package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
)

type WorkshopsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *WorkshopsResource) List(c buffalo.Context) error {

	workshops := models.GetWorkshops()
	c.Set("workshops", workshops)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("workshoplist.html"))
}

// Show default implementation. note that we're listening
// at /speakers but actually getting presentations.  This is important if your
// don't want to be confused at some future point.
func (v *WorkshopsResource) Show(c buffalo.Context) error {
	ids := c.Param("speaker_ID")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.Error(400, err)
	}
	s, err := models.GetFullWorkshop(id)
	if err != nil {
		return c.Error(404, err)
	}

	c.Set("w", s)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("partials/_workshop.html"))
}
