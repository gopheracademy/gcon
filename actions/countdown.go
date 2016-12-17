package actions

import "github.com/markbates/buffalo"

// CountdownHandler serves the countdown page.
func CountdownHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("countdown.html"))
}
