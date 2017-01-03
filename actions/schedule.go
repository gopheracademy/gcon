package actions

import "github.com/gobuffalo/buffalo"

type SponsorsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *SponsorsResource) List(c buffalo.Context) error {
	return c.Render(200, publicR.HTML("soon.html"))

}

// Show default implementation.
func (v *SponsorsResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Sponsors#Show"))
}
