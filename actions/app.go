package actions

import (
	"net/http"

	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
	"github.com/gopheracademy/gcon/models"
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
