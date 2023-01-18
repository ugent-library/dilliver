package controllers

import (
	"errors"
	"net/http"

	"github.com/ugent-library/deliver/httperror"
	"github.com/ugent-library/deliver/models"
	"github.com/ugent-library/deliver/view"
)

type Errors struct {
	forbiddenView view.View
	notFoundView  view.View
}

func NewErrors() *Errors {
	return &Errors{
		forbiddenView: view.MustNew("share/page", "errors/forbidden").Status(403),
		notFoundView:  view.MustNew("share/page", "errors/not_found").Status(404),
	}
}

func (h *Errors) Forbidden(c *Ctx) error {
	return h.forbiddenView.Render(c.Res, c)
}

func (h *Errors) NotFound(c *Ctx) error {
	return h.notFoundView.Render(c.Res, c)
}

func (h *Errors) HandleError(c *Ctx, err error) {
	if err == models.ErrNotFound {
		err = httperror.NotFound
	}

	var httpErr *httperror.Error
	if !errors.As(err, &httpErr) {
		httpErr = httperror.InternalServerError
	}

	switch httpErr.Code {
	case http.StatusUnauthorized:
		c.RedirectTo("login")
	case http.StatusForbidden:
		if err := h.Forbidden(c); err != nil {
			h.HandleError(c, err)
		}
	case http.StatusNotFound:
		if err := h.NotFound(c); err != nil {
			h.HandleError(c, err)
		}
	default:
		c.Log.Error(err)
		http.Error(c.Res, http.StatusText(httpErr.Code), httpErr.Code)
	}
}
