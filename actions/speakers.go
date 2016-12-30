package actions

import "github.com/gobuffalo/buffalo"

type SpeakersResource struct {
	buffalo.Resource
}

func init() {
	App().Resource("/speakers", &SpeakersResource{&buffalo.BaseResource{}})
}

// List default implementation.
func (v *SpeakersResource) List(c buffalo.Context) error {
	return c.Render(200, publicR.String("Speakers#List"))
}

// Show default implementation.
func (v *SpeakersResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Speakers#Show"))
}
