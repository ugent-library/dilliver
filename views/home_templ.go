// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ugent-library/deliver/ctx"

func Home(c *ctx.Ctx) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-100 u-scroll-wrapper\"><div class=\"u-scroll-wrapper__body p-8\"><div class=\"bc-toolbar bc-toolbar--auto align-items-start\"><div class=\"bc-toolbar-left\"><div class=\"bc-toolbar-item\"><h1>Deliver</h1><p class=\"c-intro\">Supporting library services for UGent Librarians.</p></div></div><div class=\"bc-toolbar-right\"><div class=\"bc-toolbar-item\"><a class=\"btn btn-primary\" href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 templ.SafeURL = templ.URL(c.Path("login").String())
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var3)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><i class=\"if if-arrow-right mt-0 ms-2\"></i> <span class=\"btn-text me-2\">Log in</span></a></div></div></div><div class=\"row mt-8\"><div class=\"col-12 col-lg-6 mb-6\"><div class=\"card h-100\"><div class=\"card-body\"><div class=\"bc-avatar-and-text mb-6\"><div class=\"bc-avatar bc-avatar--medium bc-avatar--rounded bc-avatar--light-blue\"><i class=\"if if-copy\"></i></div><div class=\"bc-avatar-text\"><h3>Share folders</h3></div></div><p class=\"text-muted\">Create temporary folders to share library specific documents with the library public.</p></div></div></div><div class=\"col-12 col-lg-6 mb-6\"><div class=\"card h-100\"><div class=\"card-body\"><div class=\"bc-avatar-and-text mb-6\"><div class=\"bc-avatar bc-avatar--medium bc-avatar--rounded bc-avatar--light-blue\"><i class=\"if if-file\"></i></div><div class=\"bc-avatar-text\"><h3>Upload &amp; monitor documents</h3></div></div><p class=\"text-muted\">Upload any type of document, up to 2GB. See whether documents have been downloaded already.</p></div></div></div></div><div class=\"row\"><div class=\"col-lg-6 mb-6\"><div class=\"card h-100\"><div class=\"card-body\"><div class=\"bc-avatar-and-text mb-6\"><div class=\"bc-avatar bc-avatar--medium bc-avatar--rounded bc-avatar--light-blue\"><i class=\"if if-edit\"></i></div><div class=\"bc-avatar-text\"><h3>Manage folder content &amp; acces</h3></div></div><p class=\"text-muted\">Swap out documents and adapt expiration dates for folders whenever you need to.</p></div></div></div></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = pageLayout(c, "Deliver").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
