package actions

import (
	"net/http"

	"github.com/gopheracademy/gcon/actions/admin"
	"github.com/gopheracademy/gcon/models"
	"github.com/kiasaki/gcon/actions/countdown"
	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: "development",
	})

	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())

	a.GET("/", countdown.CountdownHandler)

	adm := a.Group("/admin")
	adm.GET("/", admin.AdminHandler)

	return a
}
