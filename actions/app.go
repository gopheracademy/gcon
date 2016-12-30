package actions

import (
	"log"
	"os"
	"path/filepath"
	"strings"

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
	at, err := filepath.Abs("../assets")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Environment:", ENV)
	log.Println("Assets:", strings.Join(strings.Split(at, "/")[5:], "/"))
	a.Use(middleware.PopTransaction(models.DB))
	a.ServeFiles("/assets", assetsPath())
	a.GET("/home", HomeHandler)
	a.GET("/", CountdownHandler)
	adm := a.Group("/admin")
	adm.GET("/index", AdminHandler)

	return a
}
