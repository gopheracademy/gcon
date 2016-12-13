package actions

import "github.com/markbates/buffalo"

// AboutHandler serves the about page
func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about.html"))
}
