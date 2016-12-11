package actions

import (
	"net/http"

	"github.com/gopheracademy/gcon/models"
	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
)

func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: "development",
	})

	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)

	return a
}
