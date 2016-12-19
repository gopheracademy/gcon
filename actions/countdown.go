package actions

import "github.com/markbates/buffalo"

// CountdownHandler serves the countdown page.
func CountdownHandler(c buffalo.Context) error {
	r.HTMLLayout = "/public/countdown.html"
	return c.Render(200, r.HTML("/public/index.html"))
}
