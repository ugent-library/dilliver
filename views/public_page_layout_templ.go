// Code generated by templ@v0.2.334 DO NOT EDIT.

package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/ugent-library/deliver/ctx"

func publicPageLayout(c *ctx.Ctx, title string, content templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<!doctype html><html class=\"u-maximize-height\" dir=\"ltr\" lang=\"en\"><head><meta charset=\"utf-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta name=\"robots\" content=\"noindex\"><link rel=\"stylesheet\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(c.AssetPath("/css/app.css")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><link rel=\"icon\" href=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(c.AssetPath("/favicon.ico")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><script type=\"application/javascript\" src=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(c.AssetPath("/js/app.js")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_2 := ``
		_, err = templBuffer.WriteString(var_2)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</script><title>")
		if err != nil {
			return err
		}
		var var_3 string = title
		_, err = templBuffer.WriteString(templ.EscapeString(var_3))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</title></head><body class=\"u-maximize-height overflow-hidden u-scroll-wrapper\"><header>")
		if err != nil {
			return err
		}
		err = pageBanner(c).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("<div class=\"bc-navbar bc-navbar--small bc-navbar--bordered-bottom bc-navbar--white bc-navbar--fixed bc-navbar--scrollable shadow-sm px-4\"><div class=\"bc-toolbar bc-toolbar-sm\"><div class=\"bc-toolbar-left\"><div class=\"bc-toolbar-item\"><nav aria-label=\"breadcrumb\"><ol class=\"breadcrumb\"><li class=\"breadcrumb-item\"><a href=\"")
		if err != nil {
			return err
		}
		var var_4 templ.SafeURL = templ.URL(c.PathTo("home").String())
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_4)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\"><img class=\"d-none d-lg-inline-block\" src=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(c.AssetPath("/images/ghent-university-library-logo.svg")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" alt=\"Ghent University Library\"><img class=\"d-inline-block d-lg-none\" src=\"")
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString(templ.EscapeString(c.AssetPath("/images/ghent-university-library-mark.svg")))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\" alt=\"Ghent University Library\"></a></li><li class=\"breadcrumb-item\" aria-current=\"page\"><a href=\"")
		if err != nil {
			return err
		}
		var var_5 templ.SafeURL = templ.URL(c.PathTo("home").String())
		_, err = templBuffer.WriteString(templ.EscapeString(string(var_5)))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("\">")
		if err != nil {
			return err
		}
		var_6 := `Home`
		_, err = templBuffer.WriteString(var_6)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</a></li></ol></nav></div></div><div class=\"bc-toolbar-right\"><div class=\"bc-toolbar-item\"><div id=\"side-panels\"><ul class=\"nav nav-main\"></ul></div></div></div></div></div></header><main>")
		if err != nil {
			return err
		}
		err = content.Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</main></body></html>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
