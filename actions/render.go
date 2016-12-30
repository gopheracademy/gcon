package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/render/resolvers"
)

// renderer for Admin Pages
var adminR *render.Engine

// renderer for Public Pages
var publicR *render.Engine

var pResolver = &resolvers.RiceBox{
	Box: rice.MustFindBox("../templates/public"),
}

var aResolver = &resolvers.RiceBox{
	Box: rice.MustFindBox("../templates/admin"),
}

func init() {
	adminR = render.New(render.Options{
		HTMLLayout:     "/admin/main.html",
		CacheTemplates: ENV == "production",
		FileResolver:   aResolver,
	})

	publicR = render.New(render.Options{
		HTMLLayout:     "/public/main.html",
		CacheTemplates: ENV == "production",
		FileResolver:   pResolver,
	})
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../assets")
	return box.HTTPBox()
}
