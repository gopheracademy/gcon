package countdown

import "github.com/markbates/buffalo"

func CountdownHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}
