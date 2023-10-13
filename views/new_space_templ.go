// Code generated by templ@v0.2.364 DO NOT EDIT.

package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/ugent-library/deliver/ctx"
	"github.com/ugent-library/deliver/models"
	"github.com/ugent-library/deliver/validate"
	"strings"
)

func NewSpace(c *ctx.Ctx, space *models.Space, errs *validate.Errors) templ.Component {
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
		var_2 := templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
			templBuffer, templIsBuffer := w.(*bytes.Buffer)
			if !templIsBuffer {
				templBuffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templBuffer)
			}
			_, err = templBuffer.WriteString("<div class=\"w-100 u-scroll-wrapper\"><div class=\"bg-white\"><div class=\"bc-navbar bc-navbar--xlarge bc-navbar--white bc-navbar--bordered-bottom\"><div class=\"bc-toolbar\"><div class=\"bc-toolbar-left\"><div class=\"bc-toolbar-item\"><h4 class=\"bc-toolbar-title\">")
			if err != nil {
				return err
			}
			var_3 := `Make a new space`
			_, err = templBuffer.WriteString(var_3)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</h4></div></div></div></div></div><div class=\"u-scroll-wrapper__body p-6\"><div class=\"card w-100 mb-6\"><div class=\"card-header\"><div class=\"bc-toolbar\"><div class=\"bc-toolbar-left\"><div class=\"bc-toolbar-item\">")
			if err != nil {
				return err
			}
			var_4 := `Make a new space`
			_, err = templBuffer.WriteString(var_4)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</div></div><div class=\"bc-toolbar-right\"><div class=\"bc-toolbar-item\"><a class=\"btn btn-link btn-link-muted\" href=\"")
			if err != nil {
				return err
			}
			var var_5 templ.SafeURL = templ.URL(c.PathTo("spaces").String())
			_, err = templBuffer.WriteString(templ.EscapeString(string(var_5)))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\"><i class=\"if if-close\"></i><span class=\"btn-text\">")
			if err != nil {
				return err
			}
			var_6 := `Cancel`
			_, err = templBuffer.WriteString(var_6)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span></a></div><div class=\"bc-toolbar-item\"><button class=\"btn btn-primary\" data-submit-target=\"#create-space\"><i class=\"if if-check\"></i><span class=\"btn-text\">")
			if err != nil {
				return err
			}
			var_7 := `Make Space`
			_, err = templBuffer.WriteString(var_7)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</span></button></div></div></div></div><div class=\"card-body\"><form action=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(c.PathTo("createSpace").String()))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\" method=\"POST\" id=\"create-space\">")
			if err != nil {
				return err
			}
			err = csrfField(c).Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("<div class=\"row mb-6\"><label class=\"col-lg-3 col-xl-2 col-form-label\" for=\"space-name\">")
			if err != nil {
				return err
			}
			var_8 := `Space name`
			_, err = templBuffer.WriteString(var_8)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><div class=\"col-lg-5 col-xl-4\">")
			if err != nil {
				return err
			}
			if e := errs.Get("name"); e != nil {
				_, err = templBuffer.WriteString("<input class=\"form-control is-invalid\" type=\"text\" value=\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(templ.EscapeString(space.Name))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\" id=\"space-name\" name=\"name\" aria-invalid=\"true\" aria-describedby=\"space-name-invalid\"> <small class=\"invalid-feedback\" id=\"space-name-invalid\">")
				if err != nil {
					return err
				}
				var var_9 string = e.Error()
				_, err = templBuffer.WriteString(templ.EscapeString(var_9))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("</small>")
				if err != nil {
					return err
				}
			} else {
				_, err = templBuffer.WriteString("<input class=\"form-control\" type=\"text\" value=\"")
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString(templ.EscapeString(space.Name))
				if err != nil {
					return err
				}
				_, err = templBuffer.WriteString("\" id=\"space-name\" name=\"name\">")
				if err != nil {
					return err
				}
			}
			_, err = templBuffer.WriteString("</div></div><div class=\"row\"><label class=\"col-lg-3 col-xl-2 col-form-label\" for=\"space-admins\">")
			if err != nil {
				return err
			}
			var_10 := `Space admins`
			_, err = templBuffer.WriteString(var_10)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</label><div class=\"col-lg-5 col-xl-4\"><input class=\"form-control\" type=\"text\" value=\"")
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString(templ.EscapeString(strings.Join(space.Admins, ",")))
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("\" id=\"space-admins\" name=\"admins\"><p class=\"small form-text text-muted\">")
			if err != nil {
				return err
			}
			var_11 := `Separate usernames with a comma.`
			_, err = templBuffer.WriteString(var_11)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</p></div></div></form></div></div></div></div>")
			if err != nil {
				return err
			}
			if !templIsBuffer {
				_, err = io.Copy(w, templBuffer)
			}
			return err
		})
		err = pageLayout(c, "New space").Render(templ.WithChildren(ctx, var_2), templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
