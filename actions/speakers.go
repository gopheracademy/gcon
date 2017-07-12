package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
)

type SpeakersResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *SpeakersResource) List(c buffalo.Context) error {

	presentations := models.GetPresentations()
	c.Set("presentations", presentations)

	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("speakerlist.html"))
}

// Show default implementation. note that we're listening
// at /speakers but actually getting presentations.  This is important if your
// don't want to be confused at some future point.
func (v *SpeakersResource) Show(c buffalo.Context) error {
	ids := c.Param("speaker_ID")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.Error(400, err)
	}
	s, err := models.GetFullPresentation(id)
	if err != nil {
		return c.Error(404, err)
	}

	c.Set("p", s)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("speaker.html"))
}
