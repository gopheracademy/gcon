package actions

import (
	"log"
	"os"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/middleware"
	"github.com/gopheracademy/gcon/models"
	"github.com/markbates/going/defaults"
)

// ENV is the environment, computed below from environment variable or defaulted
// to development
var ENV = defaults.String(os.Getenv("GO_ENV"), "development")

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() *buffalo.App {
	a := buffalo.Automatic(buffalo.Options{
		Env: ENV,
	})
	log.Println("Environment:", ENV)
	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", HomeHandler)
	a.Resource("/speakers", &SpeakersResource{&buffalo.BaseResource{}})
	a.Resource("/schedule", &ScheduleResource{&buffalo.BaseResource{}})
	a.Resource("/hotels", &HotelsResource{&buffalo.BaseResource{}})
	a.Resource("/sponsors", &SponsorsResource{&buffalo.BaseResource{}})
	adm := a.Group("/admin")
	adm.GET("/index", AdminHandler)

	return a
}
