package actions

import "github.com/gobuffalo/buffalo"

type SpeakersResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *SpeakersResource) List(c buffalo.Context) error {
	return c.Render(200, publicR.HTML("soon.html"))
}

// Show default implementation.
func (v *SpeakersResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Speakers#Show"))
}
