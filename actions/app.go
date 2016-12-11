package actions

import (
	"net/http"

	"github.com/gopheracademy/gcon/models"
	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
)

// App is the default web handler for the application
// Add Routes and middleware here
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: "development",
	})

	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)

	return a
}
