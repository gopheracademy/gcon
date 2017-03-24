package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

// adminR is the renderer for Admin Pages
var adminR *render.Engine

// publicR is the renderer for Public Pages
var publicR *render.Engine

func init() {
	adminR = render.New(render.Options{
		HTMLLayout: "main.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates/admin"),
		Helpers: map[string]interface{}{
			"rightleft": RightLeft,
		},
	})

	publicR = render.New(render.Options{
		HTMLLayout:   "main.html",
		TemplatesBox: packr.NewBox("../templates/public"),
		Helpers: map[string]interface{}{
			"rightleft": RightLeft,
			"offset":    Offset,
		},
	})
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../assets")
	return box.HTTPBox()
}
func Offset(number int) string {
	if number%2 == 0 {
		return ""
	}
	return "col-sm-offset-6"
}

func RightLeft(number int) string {
	if number%2 == 0 {
		return "right"
	}
	return "left"
}
