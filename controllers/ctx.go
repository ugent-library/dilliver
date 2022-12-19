package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/ugent-library/dilliver/models"
	"go.uber.org/zap"
)

const (
	Info = "info"

	flashSessionKey = "flash"
)

func Wrapper(c Config) func(func(http.ResponseWriter, *http.Request, Ctx) error) http.HandlerFunc {
	return func(fn func(http.ResponseWriter, *http.Request, Ctx) error) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ctx := Ctx{
				Config:    c,
				path:      mux.Vars(r),
				CSRFToken: csrf.Token(r),
				CSRFTag:   csrf.TemplateField(r),
			}
			if err := ctx.loadSession(w, r); err != nil {
				ctx.handleError(w, r, err)
				return
			}
			if err := fn(w, r, ctx); err != nil {
				ctx.handleError(w, r, err)
			}
		}
	}
}

type Config struct {
	Log          *zap.SugaredLogger
	SessionStore sessions.Store
	SessionName  string
	Router       *mux.Router
}

type Flash struct {
	Type string
	Body template.HTML
}

type Var map[string]any

type Ctx struct {
	Config
	path      map[string]string
	CSRFToken string
	CSRFTag   template.HTML
	Flash     []Flash
	Var       Var
}

func (c Ctx) Path(k string) string {
	return c.path[k]
}

func (c Ctx) Yield(v Var) Ctx {
	c.Var = v
	return c
}

func (c Ctx) URL(route string, pairs ...string) *url.URL {
	r := c.Router.Get(route)
	if r == nil {
		panic(fmt.Errorf("unknown route '%s'", route))
	}
	u, err := r.URL(pairs...)
	if err != nil {
		panic(fmt.Errorf("can't reverse route '%s': %w", route, err))
	}
	return u
}

func (c Ctx) URLPath(route string, pairs ...string) *url.URL {
	r := c.Router.Get(route)
	if r == nil {
		panic(fmt.Errorf("unknown route '%s'", route))
	}
	u, err := r.URLPath(pairs...)
	if err != nil {
		panic(fmt.Errorf("can't reverse route '%s': %w", route, err))
	}
	return u
}

func (c Ctx) RedirectTo(w http.ResponseWriter, r *http.Request, route string, pairs ...string) {
	http.Redirect(w, r, c.URLPath(route, pairs...).String(), http.StatusSeeOther)
}

func (c Ctx) PersistFlash(w http.ResponseWriter, r *http.Request, f Flash) error {
	s, err := c.SessionStore.Get(r, c.SessionName)
	if err != nil {
		return fmt.Errorf("couldn't get session data: %w", err)
	}

	s.AddFlash(f, flashSessionKey)

	if err := s.Save(r, w); err != nil {
		return fmt.Errorf("couldn't save session data: %w", err)
	}
	return nil
}

func (c *Ctx) loadSession(w http.ResponseWriter, r *http.Request) error {
	s, err := c.SessionStore.Get(r, c.SessionName)
	if err != nil {
		return fmt.Errorf("couldn't get session data: %w", err)
	}

	flashes := s.Flashes(flashSessionKey)

	if err := s.Save(r, w); err != nil {
		return fmt.Errorf("couldn't save session data: %w", err)
	}

	for _, f := range flashes {
		c.Flash = append(c.Flash, f.(Flash))
	}

	return nil
}

func (c *Ctx) handleError(w http.ResponseWriter, r *http.Request, err error) {
	if errors.Is(err, models.ErrNotFound) {
		c.Router.NotFoundHandler.ServeHTTP(w, r)
		return
	}

	// TODO error page
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
