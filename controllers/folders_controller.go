package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/ugent-library/bind"
	"github.com/ugent-library/deliver/ctx"
	"github.com/ugent-library/deliver/htmx"
	"github.com/ugent-library/deliver/models"
	"github.com/ugent-library/deliver/objectstore"
	"github.com/ugent-library/deliver/repositories"
	"github.com/ugent-library/deliver/validate"
	"github.com/ugent-library/deliver/views"
	"github.com/ugent-library/httperror"
	"github.com/ugent-library/httpx"
)

type FoldersController struct {
	repo        *repositories.Repo
	storage     objectstore.Store
	maxFileSize int64
}

type FolderForm struct {
	Name string `form:"name"`
}

func NewFoldersController(r *repositories.Repo, s objectstore.Store, maxFileSize int64) *FoldersController {
	return &FoldersController{
		repo:        r,
		storage:     s,
		maxFileSize: maxFileSize,
	}
}

func (h *FoldersController) Show(w http.ResponseWriter, r *http.Request) {
	c := ctx.Get(r.Context())

	folderID := c.RouteParam("folderID")
	folder, err := h.repo.Folders.Get(r.Context(), folderID)
	if err != nil {
		c.HandleError(err)
		return
	}

	if htmx.Request(r) {
		httpx.RenderHTML(w, http.StatusOK, views.Files(c, folder.Files))
		return
	}

	httpx.RenderHTML(w, http.StatusOK, views.Page(c, &views.ShowFolder{
		Folder:      folder,
		MaxFileSize: h.maxFileSize,
	}))
}

func (h *FoldersController) Edit(w http.ResponseWriter, r *http.Request) {
	c := ctx.Get(r.Context())

	folderID := c.RouteParam("folderID")

	folder, err := h.repo.Folders.Get(r.Context(), folderID)
	if err != nil {
		c.HandleError(err)
		return
	}

	if !c.IsSpaceAdmin(c.User, folder.Space) {
		c.HandleError(httperror.Forbidden)
		return
	}

	httpx.RenderHTML(w, http.StatusOK, views.Page(c, &views.EditFolder{
		Folder:           folder,
		ValidationErrors: validate.NewErrors(),
	}))
}

func (h *FoldersController) Update(w http.ResponseWriter, r *http.Request) {
	c := ctx.Get(r.Context())

	folderID := c.RouteParam("folderID")

	folder, err := h.repo.Folders.Get(r.Context(), folderID)
	if err != nil {
		c.HandleError(err)
		return
	}

	if !c.IsSpaceAdmin(c.User, folder.Space) {
		c.HandleError(httperror.Forbidden)
		return
	}

	b := FolderForm{}
	if err := bind.Form(r, &b); err != nil {
		c.HandleError(errors.Join(httperror.BadRequest, err))
		return
	}

	folder.Name = b.Name

	if err := h.repo.Folders.Update(r.Context(), folder); err != nil {
		validationErrors := validate.NewErrors()
		if err != nil && !errors.As(err, &validationErrors) {
			c.HandleError(err)
			return
		}

		httpx.RenderHTML(w, http.StatusOK, views.Page(c, &views.EditFolder{
			Folder:           folder,
			ValidationErrors: validationErrors,
		}))
		return
	}

	loc := c.PathTo("folder", "folderID", folder.ID).String()
	http.Redirect(w, r, loc, http.StatusSeeOther)
}

func (h *FoldersController) Delete(w http.ResponseWriter, r *http.Request) {
	c := ctx.Get(r.Context())

	folderID := c.RouteParam("folderID")

	folder, err := h.repo.Folders.Get(r.Context(), folderID)
	if err != nil {
		c.HandleError(err)
		return
	}

	if !c.IsSpaceAdmin(c.User, folder.Space) {
		c.HandleError(httperror.Forbidden)
		return
	}

	if err := h.repo.Folders.Delete(r.Context(), folderID); err != nil {
		c.HandleError(err)
		return
	}

	c.Hub.Send("space."+folder.Space.ID, views.AddFlash(ctx.Flash{
		Type: "info",
		Body: fmt.Sprintf("%s just deleted the folder %s.", c.User.Name, folder.Name),
	}))
	c.Hub.Send("folder."+folder.ID, views.AddFlash(ctx.Flash{
		Type: "error",
		Body: fmt.Sprintf("%s just deleted this folder.", c.User.Name),
	}))
	c.PersistFlash(ctx.Flash{
		Type:         "info",
		Body:         "Folder deleted succesfully",
		DismissAfter: 3 * time.Second,
	})

	loc := c.PathTo("space", "spaceName", folder.Space.Name).String()
	http.Redirect(w, r, loc, http.StatusSeeOther)
}

func (h *FoldersController) UploadFile(w http.ResponseWriter, r *http.Request) {
	c := ctx.Get(r.Context())

	folderID := c.RouteParam("folderID")

	folder, err := h.repo.Folders.Get(r.Context(), folderID)
	if err != nil {
		c.HandleError(httperror.NotFound)
		return
	}

	if !c.IsSpaceAdmin(c.User, folder.Space) {
		c.HandleError(httperror.Forbidden)
		return
	}

	// TODO: retrieve content type by content sniffing without interfering with streaming body
	contentLength, _ := strconv.ParseInt(r.Header.Get("Content-Length"), 10, 64)

	// request header only accepts ISO-8859-1 so we had to escape it
	uploadFilename, _ := url.QueryUnescape(r.Header.Get("X-Upload-Filename"))

	file := &models.File{
		FolderID:    folderID,
		ID:          ulid.Make().String(),
		Name:        uploadFilename,
		ContentType: r.Header.Get("Content-Type"),
		Size:        contentLength,
	}

	// TODO get size
	md5, err := h.storage.Add(r.Context(), file.ID, r.Body)
	if err != nil {
		c.HandleError(err)
		return
	}

	file.MD5 = md5

	if err := h.repo.Files.Create(r.Context(), file); err != nil {
		c.HandleError(err)
	}
}

func (h *FoldersController) Share(w http.ResponseWriter, r *http.Request) {
	c := ctx.Get(r.Context())

	folderID := c.RouteParam("folderID")
	folder, err := h.repo.Folders.Get(r.Context(), folderID)
	if err != nil {
		c.HandleError(err)
		return
	}

	httpx.RenderHTML(w, http.StatusOK, views.PublicPage(c, &views.ShareFolder{
		Folder: folder,
	}))
}
