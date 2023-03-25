// Code generated by qtc from "show_space.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/show_space.qtpl:1
package views

//line views/show_space.qtpl:1
import "github.com/ugent-library/friendly"

//line views/show_space.qtpl:2
import "github.com/ugent-library/deliver/ctx"

//line views/show_space.qtpl:3
import "github.com/ugent-library/deliver/models"

//line views/show_space.qtpl:4
import "github.com/ugent-library/deliver/validate"

//line views/show_space.qtpl:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/show_space.qtpl:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/show_space.qtpl:7
type ShowSpace struct {
	Space            *models.Space
	UserSpaces       []*models.Space
	Folder           *models.Folder
	ValidationErrors *validate.Errors
}

//line views/show_space.qtpl:15
func (v *ShowSpace) StreamTitle(qw422016 *qt422016.Writer, c *ctx.Ctx) {
//line views/show_space.qtpl:15
	qw422016.E().S(v.Space.Name)
//line views/show_space.qtpl:15
}

//line views/show_space.qtpl:15
func (v *ShowSpace) WriteTitle(qq422016 qtio422016.Writer, c *ctx.Ctx) {
//line views/show_space.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/show_space.qtpl:15
	v.StreamTitle(qw422016, c)
//line views/show_space.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line views/show_space.qtpl:15
}

//line views/show_space.qtpl:15
func (v *ShowSpace) Title(c *ctx.Ctx) string {
//line views/show_space.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/show_space.qtpl:15
	v.WriteTitle(qb422016, c)
//line views/show_space.qtpl:15
	qs422016 := string(qb422016.B)
//line views/show_space.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/show_space.qtpl:15
	return qs422016
//line views/show_space.qtpl:15
}

//line views/show_space.qtpl:17
func (v *ShowSpace) StreamContent(qw422016 *qt422016.Writer, c *ctx.Ctx) {
//line views/show_space.qtpl:17
	qw422016.N().S(`
    <div class="c-sub-sidebar c-sidebar--bordered">
        <div class="bc-navbar bc-navbar--xlarge bc-navbar--bordered-bottom">
            <div class="bc-toolbar">
                <div class="bc-toolbar-left">
                    <div class="bc-toolbar-item">
                        <h1 class="bc-toolbar-title">Your deliver spaces</h1>
                    </div>
                </div>
            </div>
        </div>
        <div class="c-sub-sidebar__menu my-6">
            <nav>
                <ul class="c-sub-sidebar-menu">
                    `)
//line views/show_space.qtpl:31
	for _, s := range v.UserSpaces {
//line views/show_space.qtpl:31
		qw422016.N().S(`
                    <li class="c-sub-sidebar__item`)
//line views/show_space.qtpl:32
		if s.ID == v.Space.ID {
//line views/show_space.qtpl:32
			qw422016.N().S(` c-sub-sidebar__item--active`)
//line views/show_space.qtpl:32
		}
//line views/show_space.qtpl:32
		qw422016.N().S(`">
                        <a href="`)
//line views/show_space.qtpl:33
		qw422016.E().S(c.PathTo("space", "spaceName", s.Name).String())
//line views/show_space.qtpl:33
		qw422016.N().S(`">
                            <span class="c-sidebar__label">`)
//line views/show_space.qtpl:34
		qw422016.E().S(s.Name)
//line views/show_space.qtpl:34
		qw422016.N().S(`</span>
                        </a>
                    </li>
                    `)
//line views/show_space.qtpl:37
	}
//line views/show_space.qtpl:37
	qw422016.N().S(`
                    `)
//line views/show_space.qtpl:38
	if c.IsAdmin(c.User) {
//line views/show_space.qtpl:38
		qw422016.N().S(`
                    <li class="c-sub-sidebar__item">
                        <a href="`)
//line views/show_space.qtpl:40
		qw422016.E().S(c.PathTo("new_space").String())
//line views/show_space.qtpl:40
		qw422016.N().S(`">
                            <span class="c-sidebar__label">
                                <i class="if if-add"></i>
                                Make a new space
                            </span>
                        </a>
                    </li>
                    `)
//line views/show_space.qtpl:47
	}
//line views/show_space.qtpl:47
	qw422016.N().S(`
                </ul>
            </nav>
        </div>
    </div>

    <div class="w-100 u-scroll-wrapper">
        <div class="bg-white">
            <div class="bc-navbar bc-navbar--xlarge bc-navbar--white bc-navbar--bordered-bottom">
                <div class="bc-toolbar">
                    <div class="bc-toolbar-left">
                        <div class="bc-toolbar-item">
                            <h1 class="bc-toolbar-title">`)
//line views/show_space.qtpl:59
	qw422016.E().S(v.Space.Name)
//line views/show_space.qtpl:59
	qw422016.N().S(` folders</h1>
                        </div>
                    </div>
                    `)
//line views/show_space.qtpl:62
	if c.IsAdmin(c.User) {
//line views/show_space.qtpl:62
		qw422016.N().S(`
                    <div class="bc-toolbar-right">
                        <div class="bc-toolbar-item">
                            <a class="btn btn-link btn-link-muted" href="`)
//line views/show_space.qtpl:65
		qw422016.E().S(c.PathTo("edit_space", "spaceName", v.Space.Name).String())
//line views/show_space.qtpl:65
		qw422016.N().S(`">
                                <i class="if if-edit"></i>
                                <span class="btn-text">Edit space</span>
                            </a>
                        </div>
                    </div>
                    `)
//line views/show_space.qtpl:71
	}
//line views/show_space.qtpl:71
	qw422016.N().S(`
                </div>
            </div>
        </div>
        <div class="u-scroll-wrapper__body p-6">
            <div class="card w-100 mb-6">
                <div class="card-header">
                    <div class="bc-toolbar">
                        <div class="bc-toolbar-left">
                            <div class="bc-toolbar-item">
                                <h2>Make a folder</h2>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="card-body">
                    <form action="`)
//line views/show_space.qtpl:87
	qw422016.E().S(c.PathTo("create_folder", "spaceName", v.Space.Name).String())
//line views/show_space.qtpl:87
	qw422016.N().S(`" method="POST">
                        `)
//line views/show_space.qtpl:88
	qw422016.N().S(c.CSRFTag)
//line views/show_space.qtpl:88
	qw422016.N().S(`
                        <div class="form-group`)
//line views/show_space.qtpl:89
	if v.ValidationErrors.Get("name") != nil {
//line views/show_space.qtpl:89
		qw422016.N().S(` is-invalid`)
//line views/show_space.qtpl:89
	}
//line views/show_space.qtpl:89
	qw422016.N().S(`">
                            <label class="c-label" for="folder-name">Folder name</label>
                            <div class="form-row">
                                <div class="col-md-6">
                                    `)
//line views/show_space.qtpl:93
	if err := v.ValidationErrors.Get("name"); err != nil {
//line views/show_space.qtpl:93
		qw422016.N().S(`
                                    <input class="form-control is-invalid" type="text" value="`)
//line views/show_space.qtpl:94
		qw422016.E().S(v.Folder.Name)
//line views/show_space.qtpl:94
		qw422016.N().S(`" id="folder-name" name="name" aria-invalid="true" aria-describedby="folder-name-invalid">
                                    <small class="invalid-feedback" id="folder-name-invalid">`)
//line views/show_space.qtpl:95
		qw422016.E().S(err.Error())
//line views/show_space.qtpl:95
		qw422016.N().S(`</small>
                                    `)
//line views/show_space.qtpl:96
	} else {
//line views/show_space.qtpl:96
		qw422016.N().S(`
                                    <input class="form-control" type="text" value="`)
//line views/show_space.qtpl:97
		qw422016.E().S(v.Folder.Name)
//line views/show_space.qtpl:97
		qw422016.N().S(`" id="folder-name" name="name">
                                    `)
//line views/show_space.qtpl:98
	}
//line views/show_space.qtpl:98
	qw422016.N().S(`
                                    <small class="form-text text-muted">
                                        We will generate a shareable public link for you.
                                        <br>
                                        Each folder will expire one month after creation date.
                                    </small>
                                </div>
                                <div class="col-md-3">
                                    <button class="btn btn-primary ml-4" type="submit">
                                        <i class="if if-check"></i>
                                        <span class="btn-text">Make folder</span>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </form>
                </div>
            </div>
            <div class="card w-100 mb-6">
                <div class="card-header">
                    <div class="bc-toolbar">
                        <div class="bc-toolbar-left">
                            <div class="bc-toolbar-item">
                                <h2>Folders</h2>
                            </div>
                        </div>
                        <div class="bc-toolbar-right">
                            <div class="bc-toolbar-item">
                                <p>Showing `)
//line views/show_space.qtpl:126
	qw422016.N().D(len(v.Space.Folders))
//line views/show_space.qtpl:126
	qw422016.N().S(` of `)
//line views/show_space.qtpl:126
	qw422016.N().D(len(v.Space.Folders))
//line views/show_space.qtpl:126
	qw422016.N().S(` folders</p>
                            </div>
                        </div>
                    </div>
                </div>
                `)
//line views/show_space.qtpl:131
	if len(v.Space.Folders) > 0 {
//line views/show_space.qtpl:131
		qw422016.N().S(`
                <div class="table-responsive">
                    <table class="table table-sm table-bordered">
                        <thead>
                            <tr>
                                <th class="table-col-lg-fixed table-col-sm-fixed-left text-nowrap">Folder</th>
                                <th class="text-nowrap">Public shareable link</th>
                                <th class="text-nowrap">Expires on</th>
                                <th class="text-nowrap">Documents</th>
                                <th class="text-nowrap">Created at</th>
                                <th class="text-nowrap">Updated at</th>
                                <th class="table-col-sm-fixed table-col-sm-fixed-right"></th>
                            </tr>
                        </thead>
                        <tbody>
                            `)
//line views/show_space.qtpl:146
		for _, f := range v.Space.Folders {
//line views/show_space.qtpl:146
			qw422016.N().S(`
                            <tr>
                                <td class="text-nowrap table-col-lg-fixed table-col-sm-fixed-left">
                                    <a href="`)
//line views/show_space.qtpl:149
			qw422016.E().S(c.PathTo("folder", "folderID", f.ID).String())
//line views/show_space.qtpl:149
			qw422016.N().S(`">`)
//line views/show_space.qtpl:149
			qw422016.E().S(f.Name)
//line views/show_space.qtpl:149
			qw422016.N().S(`</a>
                                </td>
                                <td class="text-nowrap">
                                    <div class="input-group" style="min-width: 375px;">
                                        <div class="input-group-prepend">
                                            <button type="button" class="btn btn-outline-secondary btn-copy-to-clipboard" data-value="`)
//line views/show_space.qtpl:154
			qw422016.E().S(c.URLTo("share_folder", "folderID", f.ID, "folderSlug", f.Slug()).String())
//line views/show_space.qtpl:154
			qw422016.N().S(`">
                                                <i class="if if-copy text-primary"></i>
                                                <span class="btn-text">Copy link</span>
                                            </button>
                                        </div>
                                        <input type="text" class="form-control input-select-text" readonly value="`)
//line views/show_space.qtpl:159
			qw422016.E().S(c.URLTo("share_folder", "folderID", f.ID, "folderSlug", f.Slug()).String())
//line views/show_space.qtpl:159
			qw422016.N().S(`" style="min-width: 250px;">
                                    </div>
                                </td>
                                <td class="text-nowrap">
                                    <p>`)
//line views/show_space.qtpl:163
			qw422016.E().S(f.ExpiresAt.Format("2006-01-02 15:04"))
//line views/show_space.qtpl:163
			qw422016.N().S(`</p>
                                </td>
                                <td class="text-nowrap">
                                    <p>`)
//line views/show_space.qtpl:166
			qw422016.N().D(f.FileCount)
//line views/show_space.qtpl:166
			qw422016.N().S(` files</p>
                                    <p class="small">`)
//line views/show_space.qtpl:167
			qw422016.E().S(friendly.Bytes(f.Size))
//line views/show_space.qtpl:167
			qw422016.N().S(`</p>
                                </td>
                                <td class="text-nowrap">
                                    <p>`)
//line views/show_space.qtpl:170
			qw422016.E().S(f.CreatedAt.Format("2006-01-02 15:04"))
//line views/show_space.qtpl:170
			qw422016.N().S(`</p>
                                </td>
                                <td class="text-nowrap">
                                    <p>`)
//line views/show_space.qtpl:173
			qw422016.E().S(f.UpdatedAt.Format("2006-01-02 15:04"))
//line views/show_space.qtpl:173
			qw422016.N().S(`</p>
                                </td>
                                <td class="table-col-sm-fixed table-col-sm-fixed-right">
                                    <div class="c-button-toolbar">
                                        <a class="btn btn-link" href="`)
//line views/show_space.qtpl:177
			qw422016.E().S(c.PathTo("folder", "folderID", f.ID).String())
//line views/show_space.qtpl:177
			qw422016.N().S(`">
                                            <i class="if if-draft"></i>
                                            <span class="btn-text">Open</span>
                                        </a>
                                    </div>
                                </td>
                            </tr>
                            `)
//line views/show_space.qtpl:184
		}
//line views/show_space.qtpl:184
		qw422016.N().S(`
                        </tbody>
                    </table>
                </div>
                `)
//line views/show_space.qtpl:188
	} else {
//line views/show_space.qtpl:188
		qw422016.N().S(`
                <div class="c-blank-slate c-blank-slate-muted">
                    <div class="bc-avatar">
                        <i class="if if-info-circle"></i>
                    </div>
                    <p>Make a folder to get started</p>
                </div>
                `)
//line views/show_space.qtpl:195
	}
//line views/show_space.qtpl:195
	qw422016.N().S(`
            </div>
        </div>
    </div>
`)
//line views/show_space.qtpl:199
}

//line views/show_space.qtpl:199
func (v *ShowSpace) WriteContent(qq422016 qtio422016.Writer, c *ctx.Ctx) {
//line views/show_space.qtpl:199
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/show_space.qtpl:199
	v.StreamContent(qw422016, c)
//line views/show_space.qtpl:199
	qt422016.ReleaseWriter(qw422016)
//line views/show_space.qtpl:199
}

//line views/show_space.qtpl:199
func (v *ShowSpace) Content(c *ctx.Ctx) string {
//line views/show_space.qtpl:199
	qb422016 := qt422016.AcquireByteBuffer()
//line views/show_space.qtpl:199
	v.WriteContent(qb422016, c)
//line views/show_space.qtpl:199
	qs422016 := string(qb422016.B)
//line views/show_space.qtpl:199
	qt422016.ReleaseByteBuffer(qb422016)
//line views/show_space.qtpl:199
	return qs422016
//line views/show_space.qtpl:199
}
