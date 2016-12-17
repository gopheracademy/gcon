package actions

import "github.com/markbates/buffalo"

// AdminHandler serves the admin root
func AdminHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
