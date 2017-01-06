package actions

import (
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gopheracademy/gcon/models"
	"github.com/markbates/going/defaults"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = defaults.String(os.Getenv("GO_ENV"), "development")
var app *buffalo.App

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	if app == nil {
		app = buffalo.Automatic(buffalo.Options{
			Env: ENV,
		})
	}
	app.Use(middleware.PopTransaction(models.DB))
	app.ServeFiles("/assets", assetsPath())
	app.GET("/", HomeHandler)
	app.Resource("/speakers", &SpeakersResource{&buffalo.BaseResource{}})
	app.Resource("/schedule", &ScheduleResource{&buffalo.BaseResource{}})
	app.Resource("/hotels", &HotelsResource{&buffalo.BaseResource{}})
	app.Resource("/sponsors", &SponsorsResource{&buffalo.BaseResource{}})
	adm := app.Group("/admin")
	adm.GET("/index", AdminHandler)

	return app
}
