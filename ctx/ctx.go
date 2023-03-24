package ctx

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/ugent-library/deliver/crumb"
	"github.com/ugent-library/deliver/models"
	"github.com/ugent-library/deliver/turbo"
	"github.com/ugent-library/mix"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

// TODO set __Host- cookie prefix in production
const (
	rememberCookie = "deliver.remember"
	flashCookie    = "deliver.flash"
)

type Ctx struct {
	Log       *zap.SugaredLogger // TODO use plain logger
	Req       *http.Request
	Res       http.ResponseWriter
	CSRFToken string
	CSRFTag   string
	Cookies   *crumb.CookieJar
	User      *models.User
	*models.Permissions
	Flash    []Flash
	Router   *mux.Router
	PathVars map[string]string
	Render   *render.Render
	Assets   mix.Manifest
	Turbo    *turbo.Hub
}

type Flash struct {
	Type         string
	Title        string
	Body         string
	DismissAfter time.Duration
}

type TemplateData struct {
	ctx       *Ctx
	CSRFToken string
	CSRFTag   string
	Data      any
}

func (t TemplateData) User() *models.User {
	return t.ctx.User
}

func (t TemplateData) Flash() []Flash {
	return t.ctx.Flash
}

func (t TemplateData) URLTo(name string, pairs ...string) *url.URL {
	return t.ctx.URLTo(name, pairs...)
}

func (t TemplateData) PathTo(name string, pairs ...string) *url.URL {
	return t.ctx.PathTo(name, pairs...)
}

func (t TemplateData) IsAdmin(user *models.User) bool {
	return t.ctx.IsAdmin(user)
}

func (t TemplateData) IsSpaceAdmin(user *models.User, space *models.Space) bool {
	return t.ctx.IsSpaceAdmin(user, space)
}

func (c *Ctx) Context() context.Context {
	return c.Req.Context()
}

func (c *Ctx) HTML(status int, body string) error {
	if hdr := c.Res.Header(); hdr.Get("Content-Type") == "" {
		hdr.Set("Content-Type", "text/html")
	}
	c.Res.WriteHeader(status)
	_, err := c.Res.Write([]byte(body))
	return err
}

// TODO deprecated
func (c *Ctx) HTMLX(status int, layout, tmpl string, data any) error {
	return c.Render.HTML(c.Res, status, tmpl, TemplateData{
		ctx:       c,
		CSRFToken: csrf.Token(c.Req),
		CSRFTag:   string(csrf.TemplateField(c.Req)),
		Data:      data,
	}, render.HTMLOptions{
		Layout: layout,
	})
}

func (c *Ctx) Path(param string) string {
	return c.PathVars[param]
}

func (c *Ctx) URLTo(name string, pairs ...string) *url.URL {
	route := c.Router.Get(name)
	if route == nil {
		panic(fmt.Errorf("unknown route '%s'", name))
	}
	u, err := route.URL(pairs...)
	if err != nil {
		panic(fmt.Errorf("can't reverse route '%s': %w", name, err))
	}
	if u.Host == "" {
		u.Host = c.Req.Host
	}
	if u.Scheme == "" {
		u.Scheme = c.Req.URL.Scheme
	}
	if u.Scheme == "" {
		u.Scheme = "http"
	}
	return u
}

func (c *Ctx) PathTo(name string, pairs ...string) *url.URL {
	route := c.Router.Get(name)
	if route == nil {
		panic(fmt.Errorf("unknown route '%s'", name))
	}
	u, err := route.URLPath(pairs...)
	if err != nil {
		panic(fmt.Errorf("can't reverse route '%s': %w", name, err))
	}
	return u
}

func (c *Ctx) RedirectTo(name string, pairs ...string) {
	route := c.Router.Get(name)
	if route == nil {
		panic(fmt.Errorf("unknown route '%s'", name))
	}
	u, err := route.URLPath(pairs...)
	if err != nil {
		panic(fmt.Errorf("can't reverse route '%s': %w", name, err))
	}
	http.Redirect(c.Res, c.Req, u.String(), http.StatusSeeOther)
}

func (c *Ctx) AssetPath(asset string) string {
	ap, err := c.Assets.AssetPath(asset)
	if err != nil {
		panic(err)
	}
	return ap
}

func (c *Ctx) AddFlash(f Flash) {
	c.Cookies.Append(flashCookie, f, time.Now().Add(3*time.Minute))
}

func (c *Ctx) TurboStreamTag(names ...string) string {
	cryptedNames, err := c.Turbo.EncryptStreamNames(names)
	if err != nil {
		c.Log.Error(err)
		return ""
	}
	src := c.URLTo("ws", "streams", string(cryptedNames))
	// TODO
	if c.Req.URL.Scheme == "https" {
		src.Scheme = "wss"
	} else {
		src.Scheme = "ws"
	}
	return `<turbo-stream-source src="` + src.String() + `"></turbo-stream-source>`
}
