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
var ENV = defaults.String(os.Getenv("GO_ENV"), "development")

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
func App() http.Handler {
	a := buffalo.Automatic(buffalo.Options{
		Env: ENV,
	})
	log.Println("Environment:", ENV)
	log.Println("Assets:", assetsPath())
	a.Use(setLayout())
	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/home", HomeHandler)
	a.GET("/", CountdownHandler)
	adm := a.Group("/admin")
	adm.GET("/index", AdminHandler)

	return a
}

func setLayout() buffalo.MiddlewareFunc {
	return func(h buffalo.Handler) buffalo.Handler {
		return func(c buffalo.Context) error {
			// something tells me this is a race and won't
			// work the way I expect it to.
			admin := strings.Contains(c.Request().URL.Path, "admin")
			if admin {
				r.HTMLLayout = "/admin/main.html"
			} else {
				r.HTMLLayout = "/public/main.html"
			}
			return h(c)
		}
	}
}
