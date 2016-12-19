package actions

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/markbates/buffalo/render"
	"github.com/markbates/buffalo/render/resolvers"
)

var r *render.Engine
var adminResolver = &resolvers.RiceBox{
	Box: rice.MustFindBox("../templates/admin"),
}

var publicResolver = &resolvers.RiceBox{
	Box: rice.MustFindBox("../templates/public"),
}

func init() {
	r = render.New(render.Options{
		HTMLLayout:     "main.html",
		CacheTemplates: ENV == "production",
		FileResolver:   publicResolver,
	})
}

func assetsPath() http.FileSystem {
	box := rice.MustFindBox("../assets")
	return box.HTTPBox()
}
