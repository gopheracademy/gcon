package actions

import "github.com/gobuffalo/buffalo"

type HotelsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *HotelsResource) List(c buffalo.Context) error {
	return c.Render(200, publicR.HTML("soon.html"))

}

// Show default implementation.
func (v *HotelsResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Hotels#Show"))
}
