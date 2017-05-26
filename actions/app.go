package actions

import (
	"fmt"
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
	app.GET("/about/{slug}", AboutHandler)
	app.GET("/events/{slug}", EventHandler)
	app.GET("/hotels", HotelsHandler)
	app.GET("/schedule", ScheduleHandler)
	app.Resource("/sponsors", &SponsorResource{&buffalo.BaseResource{}})
	app.Resource("/speakers", &SpeakersResource{&buffalo.BaseResource{}})
	app.Resource("/workshops", &WorkshopsResource{&buffalo.BaseResource{}})
	adm := app.Group("/admin")
	adm.GET("/index", AdminHandler)
	adm.GET("/login", LoginHandler)
	adm.POST("/login", AuthHandler)
	app.GET("/", HomeHandler)
	adm.Use(authMiddleware())
	adm.Middleware.Skip(authMiddleware(), AuthHandler, LoginHandler)

	app.ErrorHandlers[500] = func(status int, err error, c buffalo.Context) error {
		res := c.Response()
		res.WriteHeader(500)
		res.Write([]byte(fmt.Sprintf("Internal server error! : %s", err.Error())))
		return nil
	}

	return app
}

func authMiddleware() buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			name := c.Session().Get("username")
			if name == "" || name == nil {
				return c.Redirect(302, "/admin/login")
			}
			c.LogField("User", name)
			return h(c)
		}
	}
}
