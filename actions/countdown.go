package actions

import "github.com/gobuffalo/buffalo"

// CountdownHandler serves the countdown page.
func CountdownHandler(c buffalo.Context) error {
	publicR.HTMLLayout = "countdown.html"
	return c.Render(200, publicR.HTML("index.html"))
}
