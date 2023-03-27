// Code generated by qtc from "page.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/page.qtpl:1
package views

//line views/page.qtpl:1
import "github.com/ugent-library/deliver/ctx"

//line views/page.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/page.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/page.qtpl:4
type PageView interface {
//line views/page.qtpl:4
	Title(c *ctx.Ctx) string
//line views/page.qtpl:4
	StreamTitle(qw422016 *qt422016.Writer, c *ctx.Ctx)
//line views/page.qtpl:4
	WriteTitle(qq422016 qtio422016.Writer, c *ctx.Ctx)
//line views/page.qtpl:4
	Content(c *ctx.Ctx) string
//line views/page.qtpl:4
	StreamContent(qw422016 *qt422016.Writer, c *ctx.Ctx)
//line views/page.qtpl:4
	WriteContent(qq422016 qtio422016.Writer, c *ctx.Ctx)
//line views/page.qtpl:4
}

//line views/page.qtpl:10
func StreamPage(qw422016 *qt422016.Writer, c *ctx.Ctx, v PageView) {
//line views/page.qtpl:10
	qw422016.N().S(`
    <!DOCTYPE html>
    <html class="u-maximize-height" dir="ltr" lang="en">

    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="robots" content="noindex">
        <meta name="csrf-token" content="`)
//line views/page.qtpl:18
	qw422016.E().S(c.CSRFToken)
//line views/page.qtpl:18
	qw422016.N().S(`">
        <link rel="stylesheet" href="`)
//line views/page.qtpl:19
	qw422016.E().S(c.AssetPath("/css/app.css"))
//line views/page.qtpl:19
	qw422016.N().S(`" data-turbo-track="reload">
        <link rel="icon" href="`)
//line views/page.qtpl:20
	qw422016.E().S(c.AssetPath("/favicon.ico"))
//line views/page.qtpl:20
	qw422016.N().S(`" data-turbo-track="reload">
        <script type="application/javascript" src="`)
//line views/page.qtpl:21
	qw422016.E().S(c.AssetPath("/js/app.js"))
//line views/page.qtpl:21
	qw422016.N().S(`" data-turbo-track="reload"></script>
        <title>`)
//line views/page.qtpl:22
	qw422016.E().S(v.Title(c))
//line views/page.qtpl:22
	qw422016.N().S(`</title>
    </head>

    <body class="u-maximize-height overflow-hidden u-scroll-wrapper">
        <header>
            <div class="bc-navbar bc-navbar--small bc-navbar--bordered-bottom bc-navbar--white bc-navbar--fixed bc-navbar--scrollable shadow-sm px-4">
                <div class="bc-toolbar bc-toolbar-sm">
                    <div class="bc-toolbar-left">
                        <div class="bc-toolbar-item">
                            <nav aria-label="breadcrumb">
                                <ol class="breadcrumb">
                                    <li class="breadcrumb-item">
                                        <a href="`)
//line views/page.qtpl:34
	qw422016.E().S(c.PathTo("home").String())
//line views/page.qtpl:34
	qw422016.N().S(`">
                                            <img class="d-none d-lg-inline-block" src="`)
//line views/page.qtpl:35
	qw422016.E().S(c.AssetPath("/images/ghent-university-library-logo.svg"))
//line views/page.qtpl:35
	qw422016.N().S(`" alt="Ghent University Library">
                                            <img class="d-inline-block d-lg-none" src="`)
//line views/page.qtpl:36
	qw422016.E().S(c.AssetPath("/images/ghent-university-library-mark.svg"))
//line views/page.qtpl:36
	qw422016.N().S(`" alt="Ghent University Library">
                                        </a>
                                    </li>
                                    <li class="breadcrumb-item" aria-current="page">
                                        <a href="`)
//line views/page.qtpl:40
	qw422016.E().S(c.PathTo("home").String())
//line views/page.qtpl:40
	qw422016.N().S(`" class="text-muted">Home</a>
                                    </li>
                                </ol>
                            </nav>
                        </div>
                    </div>

                    <div class="bc-toolbar-right">
                        <div class="bc-toolbar-item">
                            <div id="side-panels">
                                <ul class="nav nav-main">
                                    `)
//line views/page.qtpl:51
	if c.User != nil {
//line views/page.qtpl:51
		qw422016.N().S(`
                                    <li class="nav-item">
                                        <a class="nav-link" href="https://forms.office.com/e/6D3PEnpV9M" target="_blank">
                                            <i class="if if-service"></i>
                                            <span class="btn-text">Geef feedback</span>
                                        </a>
                                    </li>
                                    `)
//line views/page.qtpl:58
	}
//line views/page.qtpl:58
	qw422016.N().S(`
                                    <li class="nav-item">
                                        <a class="nav-link" href="https://www.ugent.be/intranet/nl/op-het-werk/bibliotheek/publieksdiensten/deliverhandleiding" target="_blank">
                                            <i class="if if-book"></i>
                                            <span class="btn-text">Handleiding</span>
                                        </a>
                                    </li>
                                    <li class="nav-item">
                                        `)
//line views/page.qtpl:66
	if c.User != nil {
//line views/page.qtpl:66
		qw422016.N().S(`
                                        <div class="dropdown position-static">
                                            <button class="nav-link dropdown-toggle" role="button" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                                <div class="bc-avatar-and-text">
                                                    <div class="bc-avatar bc-avatar--sm bc-avatar--small mr-2">
                                                        <i class="if if-user"></i>
                                                    </div>
                                                    <div class="bc-avatar-text">`)
//line views/page.qtpl:73
		qw422016.E().S(c.User.Name)
//line views/page.qtpl:73
		qw422016.N().S(`</div>
                                                </div>
                                            </button>
                                            <div class="dropdown-menu mt-8">
                                                <div class="bc-avatar-and-text m-4">
                                                    <div class="bc-avatar bc-avatar--sm">
                                                        <i class="if if-user"></i>
                                                    </div>
                                                    <div class="bc-avatar-text">
                                                        <h4>`)
//line views/page.qtpl:82
		qw422016.E().S(c.User.Name)
//line views/page.qtpl:82
		qw422016.N().S(`</h4>
                                                        <p class="text-muted c-body-small">`)
//line views/page.qtpl:83
		qw422016.E().S(c.User.Email)
//line views/page.qtpl:83
		qw422016.N().S(`</p>
                                                    </div>
                                                </div>
                                                <hr class="dropdown-divider">
                                                <a class="dropdown-item" href="`)
//line views/page.qtpl:87
		qw422016.E().S(c.PathTo("logout").String())
//line views/page.qtpl:87
		qw422016.N().S(`">
                                                    <i class="if if-log-out"></i>
                                                    <span>Log out</span>
                                                </a>
                                            </div>
                                        </div>
                                        `)
//line views/page.qtpl:93
	} else {
//line views/page.qtpl:93
		qw422016.N().S(`
                                        <a class="btn btn-link btn-sm" href="`)
//line views/page.qtpl:94
		qw422016.E().S(c.PathTo("login").String())
//line views/page.qtpl:94
		qw422016.N().S(`">
                                            <i class="if if-arrow-right mt-0 ml-2"></i>
                                            <span class="btn-text mr-2">Log in</span>
                                        </a>
                                        `)
//line views/page.qtpl:98
	}
//line views/page.qtpl:98
	qw422016.N().S(`
                                    </li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </header>
        <main>
            <div class="d-flex u-maximize-height">
                <div class="c-sidebar`)
//line views/page.qtpl:109
	if c.User != nil {
//line views/page.qtpl:109
		qw422016.N().S(` c-sidebar--dark-gray`)
//line views/page.qtpl:109
	}
//line views/page.qtpl:109
	qw422016.N().S(` d-none d-lg-flex">
                    <div class="c-sidebar__menu">
                        <nav>
                            <ul class="c-sidebar-menu">
                                <li class="c-sidebar__item c-sidebar__item--active">
                                    <a href="`)
//line views/page.qtpl:114
	qw422016.E().S(c.PathTo("home").String())
//line views/page.qtpl:114
	qw422016.N().S(`">
                                        <span class="c-sidebar__icon">
                                            <i class="if if-file"></i>
                                        </span>
                                        <span class="c-sidebar__label">Deliver</span>
                                    </a>
                                </li>
                            </ul>
                        </nav>
                    </div>
                    <div class="c-sidebar__bottom">
                        <img src="`)
//line views/page.qtpl:125
	qw422016.E().S(c.AssetPath("/images/logo-ugent-white.svg"))
//line views/page.qtpl:125
	qw422016.N().S(`" alt="Logo UGent" height="48px" width="auto">
                    </div>
                </div>

                `)
//line views/page.qtpl:129
	v.StreamContent(qw422016, c)
//line views/page.qtpl:129
	qw422016.N().S(`
            </div>
        </main>

        <div id="flash-messages">
            `)
//line views/page.qtpl:134
	for _, f := range c.Flash {
//line views/page.qtpl:134
		qw422016.N().S(`
            <div class="toast" role="alert" aria-live="assertive" aria-atomic="true">
                <div class="toast-body">
                    `)
//line views/page.qtpl:137
		switch f.Type {
//line views/page.qtpl:138
		case "success":
//line views/page.qtpl:138
			qw422016.N().S(`
                    <i class="if if--success if-check-circle-fill"></i>
                    `)
//line views/page.qtpl:140
		case "info":
//line views/page.qtpl:140
			qw422016.N().S(`
                    <i class="if if--primary if-info-circle-filled"></i>
                    `)
//line views/page.qtpl:142
		case "warning":
//line views/page.qtpl:142
			qw422016.N().S(`
                    <i class="if if--warning if-alert-fill"></i>
                    `)
//line views/page.qtpl:144
		case "error":
//line views/page.qtpl:144
			qw422016.N().S(`
                    <i class="if if--error if-error-circle-fill"></i>
                    `)
//line views/page.qtpl:146
		}
//line views/page.qtpl:146
		qw422016.N().S(`
                    <div class="toast-content">
                        `)
//line views/page.qtpl:148
		if f.Title != "" {
//line views/page.qtpl:148
			qw422016.N().S(`
                        <h3 class="alert-title">`)
//line views/page.qtpl:149
			qw422016.E().S(f.Title)
//line views/page.qtpl:149
			qw422016.N().S(`</h3>
                        `)
//line views/page.qtpl:150
		}
//line views/page.qtpl:150
		qw422016.N().S(`
                        `)
//line views/page.qtpl:151
		qw422016.E().S(f.Body)
//line views/page.qtpl:151
		qw422016.N().S(`
                    </div>
                    <button class="close" type="button" aria-label="Close"
                        data-dismiss="toast"
                        `)
//line views/page.qtpl:155
		if f.DismissAfter != 0 {
//line views/page.qtpl:155
			qw422016.N().S(`
                        data-delay="`)
//line views/page.qtpl:156
			qw422016.N().DL(f.DismissAfter.Milliseconds())
//line views/page.qtpl:156
			qw422016.N().S(`"
                        `)
//line views/page.qtpl:157
		} else {
//line views/page.qtpl:157
			qw422016.N().S(`
                        data-autohide="false"
                        `)
//line views/page.qtpl:159
		}
//line views/page.qtpl:159
		qw422016.N().S(`
                    >
                        <i class="if if-close"></i>
                    </button>
                </div>
            </div>
            `)
//line views/page.qtpl:165
	}
//line views/page.qtpl:165
	qw422016.N().S(`
        </div>
    </body>
    </html>
`)
//line views/page.qtpl:169
}

//line views/page.qtpl:169
func WritePage(qq422016 qtio422016.Writer, c *ctx.Ctx, v PageView) {
//line views/page.qtpl:169
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/page.qtpl:169
	StreamPage(qw422016, c, v)
//line views/page.qtpl:169
	qt422016.ReleaseWriter(qw422016)
//line views/page.qtpl:169
}

//line views/page.qtpl:169
func Page(c *ctx.Ctx, v PageView) string {
//line views/page.qtpl:169
	qb422016 := qt422016.AcquireByteBuffer()
//line views/page.qtpl:169
	WritePage(qb422016, c, v)
//line views/page.qtpl:169
	qs422016 := string(qb422016.B)
//line views/page.qtpl:169
	qt422016.ReleaseByteBuffer(qb422016)
//line views/page.qtpl:169
	return qs422016
//line views/page.qtpl:169
}

//line views/page.qtpl:171
func StreamPublicPage(qw422016 *qt422016.Writer, c *ctx.Ctx, v PageView) {
//line views/page.qtpl:171
	qw422016.N().S(`
    <!DOCTYPE html>
    <html class="u-maximize-height" dir="ltr" lang="en">

    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <meta name="robots" content="noindex">
        <link rel="stylesheet" href="`)
//line views/page.qtpl:179
	qw422016.E().S(c.AssetPath("/css/app.css"))
//line views/page.qtpl:179
	qw422016.N().S(`" data-turbo-track="reload">
        <link rel="icon" href="`)
//line views/page.qtpl:180
	qw422016.E().S(c.AssetPath("/favicon.ico"))
//line views/page.qtpl:180
	qw422016.N().S(`" data-turbo-track="reload">
        <script type="application/javascript" src="`)
//line views/page.qtpl:181
	qw422016.E().S(c.AssetPath("/js/app.js"))
//line views/page.qtpl:181
	qw422016.N().S(`" data-turbo-track="reload"></script>
        <title>`)
//line views/page.qtpl:182
	v.StreamTitle(qw422016, c)
//line views/page.qtpl:182
	qw422016.N().S(`</title>
    </head>

    <body class="u-maximize-height overflow-hidden u-scroll-wrapper">
        <div class="u-horizontal-scroll h-100 w-100">
            <div class="u-min-w-750 h-100">
                <header>
                    <div class="bc-navbar bc-navbar--small bc-navbar--bordered-bottom bc-navbar--white bc-navbar--fixed bc-navbar--scrollable shadow-sm px-4">
                        <div class="bc-toolbar bc-toolbar-sm">
                            <div class="bc-toolbar-left">
                                <div class="bc-toolbar-item">
                                    <nav aria-label="breadcrumb">
                                        <ol class="breadcrumb">
                                            <li class="breadcrumb-item">
                                                <a href="`)
//line views/page.qtpl:196
	qw422016.E().S(c.PathTo("home").String())
//line views/page.qtpl:196
	qw422016.N().S(`">
                                                    <img class="d-none d-lg-inline-block" src="`)
//line views/page.qtpl:197
	qw422016.E().S(c.AssetPath("/images/ghent-university-library-logo.svg"))
//line views/page.qtpl:197
	qw422016.N().S(`" alt="Ghent University Library">
                                                    <img class="d-inline-block d-lg-none" src="`)
//line views/page.qtpl:198
	qw422016.E().S(c.AssetPath("/images/ghent-university-library-mark.svg"))
//line views/page.qtpl:198
	qw422016.N().S(`" alt="Ghent University Library">
                                                </a>
                                            </li>
                                            <li class="breadcrumb-item" aria-current="page">
                                                <a href="`)
//line views/page.qtpl:202
	qw422016.E().S(c.PathTo("home").String())
//line views/page.qtpl:202
	qw422016.N().S(`">Home</a>
                                            </li>
                                        </ol>
                                    </nav>
                                </div>
                            </div>
                            <div class="bc-toolbar-right">
                                <div class="bc-toolbar-item">
                                    <div id="side-panels">
                                        <ul class="nav nav-main"></ul>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </header>
                <main>
                    `)
//line views/page.qtpl:219
	v.StreamContent(qw422016, c)
//line views/page.qtpl:219
	qw422016.N().S(`
                </main>
            </div>
        </div>  
    </body>
    </html>
`)
//line views/page.qtpl:225
}

//line views/page.qtpl:225
func WritePublicPage(qq422016 qtio422016.Writer, c *ctx.Ctx, v PageView) {
//line views/page.qtpl:225
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/page.qtpl:225
	StreamPublicPage(qw422016, c, v)
//line views/page.qtpl:225
	qt422016.ReleaseWriter(qw422016)
//line views/page.qtpl:225
}

//line views/page.qtpl:225
func PublicPage(c *ctx.Ctx, v PageView) string {
//line views/page.qtpl:225
	qb422016 := qt422016.AcquireByteBuffer()
//line views/page.qtpl:225
	WritePublicPage(qb422016, c, v)
//line views/page.qtpl:225
	qs422016 := string(qb422016.B)
//line views/page.qtpl:225
	qt422016.ReleaseByteBuffer(qb422016)
//line views/page.qtpl:225
	return qs422016
//line views/page.qtpl:225
}
