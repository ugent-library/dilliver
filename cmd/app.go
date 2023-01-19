package cmd

import (
	"context"
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/ory/graceful"
	"github.com/spf13/cobra"
	"github.com/ugent-library/deliver/autosession"
	c "github.com/ugent-library/deliver/controllers"
	"github.com/ugent-library/deliver/friendly"
	"github.com/ugent-library/deliver/middleware"
	"github.com/ugent-library/deliver/mix"
	"github.com/ugent-library/deliver/models"
	"github.com/ugent-library/deliver/oidc"
	"github.com/ugent-library/deliver/ulid"
	"github.com/ugent-library/deliver/view"
	"github.com/ugent-library/deliver/zaphttp"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Start the web app server",
	Run: func(cmd *cobra.Command, args []string) {
		// setup services
		repoService, err := models.NewRepositoryService(models.RepositoryConfig{
			DB: config.DB,
		})
		if err != nil {
			logger.Fatal(err)
		}

		fileService, err := models.NewFileService(models.FileConfig{
			S3ID:     config.S3.ID,
			S3Secret: config.S3.Secret,
			S3Bucket: config.S3.Bucket,
			S3Region: config.S3.Region,
		})
		if err != nil {
			logger.Fatal(err)
		}

		// setup permissions
		permissions := &models.Permissions{
			Admins: config.Admins,
		}

		// setup auth
		oidcAuth, err := oidc.NewAuth(context.TODO(), oidc.Config{
			URL:          config.OIDC.URL,
			ClientID:     config.OIDC.ID,
			ClientSecret: config.OIDC.Secret,
			RedirectURL:  config.OIDC.RedirectURL,
			CookieName:   config.Session.Name + ".state",
			CookieSecret: []byte(config.Session.Secret),
		})
		if err != nil {
			logger.Fatal(err)
		}

		// setup assets
		assets, err := mix.New(mix.Config{
			ManifestFile: "static/mix-manifest.json",
			PublicPath:   "/static/",
		})
		if err != nil {
			logger.Fatal(err)
		}

		// setup views
		view.DefaultConfig.Funcs = template.FuncMap{
			"assetPath":     assets.AssetPath,
			"friendlyBytes": friendly.Bytes,
			"join":          strings.Join,
		}

		// setup sessions
		sessionName := config.Session.Name
		sessionStore := sessions.NewCookieStore([]byte(config.Session.Secret))
		sessionStore.MaxAge(config.Session.MaxAge)
		sessionStore.Options.Path = "/"
		sessionStore.Options.HttpOnly = true
		sessionStore.Options.Secure = config.Production
		// register types so CookieStore can serialize them
		gob.Register(&models.User{})
		gob.Register(c.Flash{})

		// setup router
		r := mux.NewRouter()
		r.StrictSlash(true)
		r.UseEncodedPath()

		// controllers
		errs := c.NewErrors()
		auth := c.NewAuth(oidcAuth)
		pages := c.NewPages()
		spaces := c.NewSpaces(repoService)
		folders := c.NewFolders(repoService, fileService)
		files := c.NewFiles(repoService, fileService)

		// request context wrapper
		wrap := c.Wrapper(c.Config{
			Router:       r,
			ErrorHandler: errs.HandleError,
			Permissions:  permissions,
		})

		// routes
		r.NotFoundHandler = wrap(errs.NotFound)
		// TODO don't apply all middleware to static file server
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
		r.Handle("/", wrap(pages.Home)).Methods("GET").Name("home")
		r.Handle("/auth/callback", wrap(auth.Callback)).Methods("GET")
		r.Handle("/logout", wrap(auth.Logout)).Methods("GET").Name("logout")
		r.Handle("/login", wrap(auth.Login)).Methods("GET").Name("login")
		r.Handle("/spaces", wrap(c.RequireUser, spaces.List)).Methods("GET").Name("spaces")
		r.Handle("/spaces/{spaceName}", wrap(c.RequireUser, spaces.Show)).Methods("GET").Name("space")
		r.Handle("/new-space", wrap(c.RequireAdmin, spaces.New)).Methods("GET").Name("new_space")
		r.Handle("/spaces", wrap(c.RequireAdmin, spaces.Create)).Methods("POST").Name("create_space")
		r.Handle("/spaces/{spaceName}/edit", wrap(c.RequireAdmin, spaces.Edit)).Methods("GET").Name("edit_space")
		r.Handle("/spaces/{spaceName}", wrap(c.RequireAdmin, spaces.Update)).Methods("PUT").Name("update_space")
		r.Handle("/spaces/{spaceName}/folders", wrap(c.RequireUser, spaces.CreateFolder)).Methods("POST").Name("create_folder")
		r.Handle("/folders/{folderID}", wrap(c.RequireUser, folders.Show)).Methods("GET").Name("folder")
		r.Handle("/folders/{folderID}/edit", wrap(c.RequireUser, folders.Edit)).Methods("GET").Name("edit_folder")
		r.Handle("/folders/{folderID}", wrap(c.RequireUser, folders.Update)).Methods("PUT").Name("update_folder")
		r.Handle("/folders/{folderID}/files", wrap(c.RequireUser, folders.UploadFile)).Methods("POST").Name("upload_file")
		r.Handle("/folders/{folderID}", wrap(c.RequireUser, folders.Delete)).Methods("DELETE").Name("delete_folder")
		r.Handle("/files/{fileID}", wrap(files.Download)).Methods("GET").Name("download_file")
		r.Handle("/files/{fileID}", wrap(c.RequireUser, files.Delete)).Methods("DELETE").Name("delete_file")
		r.Handle("/share/{folderID}:{folderSlug}", wrap(folders.Share)).Methods("GET").Name("share_folder")

		// apply these before request reaches the router
		handler := middleware.Apply(r,
			middleware.Recover(func(err any) {
				if config.Production {
					logger.With(zap.Stack("stack")).Error(err)
				} else {
					logger.Error(err)
				}
			}),
			// apply before ProxyHeaders to avoid invalid referer errors
			csrf.Protect(
				[]byte(config.Session.Secret),
				csrf.CookieName(config.Session.Name+".csrf"),
				csrf.Path("/"),
				csrf.Secure(config.Production),
				csrf.SameSite(csrf.SameSiteStrictMode),
				csrf.FieldName("csrf_token"),
			),
			handlers.HTTPMethodOverrideHandler,
			middleware.If(config.Production, handlers.ProxyHeaders),
			middleware.SetRequestID(ulid.MustGenerate),
			zaphttp.SetLogger(logger.Desugar()),
			zaphttp.LogRequests,
			autosession.Enable(autosession.GorillaSession(sessionStore, sessionName)),
		)

		// start server
		// TODO make timeouts configurable
		addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
		server := graceful.WithDefaults(&http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Minute,
			WriteTimeout: 10 * time.Minute,
		})
		logger.Infof("starting server at %s", addr)
		if err := graceful.Graceful(server.ListenAndServe, server.Shutdown); err != nil {
			logger.Fatal(err)
		}
		logger.Info("gracefully stopped server")
	},
}
