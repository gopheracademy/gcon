package actions

import "github.com/markbates/buffalo"

// HomeHandler serves the home action -- index.html
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
