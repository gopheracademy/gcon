package actions

import (
	"net/http"
	"path"
	"runtime"

	"github.com/markbates/buffalo/render"
)

var r *render.Engine

func init() {
	r = render.New(render.Options{
		TemplatesPath: templatesPath(),
		HTMLLayout:    "application.html",
	})
}

func assetsPath() http.Dir {
	if ENV == "production" {
		return http.Dir("/gcon/assets/public")
	}
	return http.Dir(fromHere("../assets/public"))
}

func templatesPath() string {
	if ENV == "production" {
		return "/gcon/templates/public"
	}
	return fromHere("../templates/public")
}

func fromHere(p string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), p)
}
