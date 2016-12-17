package actions

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gopheracademy/gcon/models"
	"github.com/markbates/buffalo"
	"github.com/markbates/buffalo/middleware"
	"github.com/markbates/going/defaults"
)

// ENV is the environment, computed below from environment variable or defaulted
// to development
var ENV string

func init() {
	ENV = defaults.String(os.Getenv("GO_ENV"), "development")
}

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: ENV,
	})
	log.Println("Environment:", ENV)
	log.Println("Assets:", assetsPath())
	a.Use(setTemplate())
	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/", CountdownHandler)
	a.GET("/home", HomeHandler)
	adm := a.Group("/admin")
	adm.GET("/index", AdminHandler)

	return a
}

func setTemplate() buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			path := "/public"
			admin := strings.Contains(c.Request().URL.Path, "admin")
			if admin {
				path = "/admin"
			}
			r.TemplatesPath = fromHere("../templates" + path)

			if ENV == "production" {
				r.TemplatesPath = "/gcon/templates" + path
			}
			// TODO: Remove after debugging
			log.Println("Template Path:", r.TemplatesPath)
			return h(c)
		}
	}
}
