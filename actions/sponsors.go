package actions

import "github.com/gobuffalo/buffalo"

type SponsorsResource struct {
	buffalo.Resource
}

func init() {
	App().Resource("/sponsors", &SponsorsResource{&buffalo.BaseResource{}})
}

// List default implementation.
func (v *SponsorsResource) List(c buffalo.Context) error {
	return c.Render(200, publicR.String("Sponsors#List"))
}

// Show default implementation.
func (v *SponsorsResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Sponsors#Show"))
}
