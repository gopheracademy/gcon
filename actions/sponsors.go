package actions

import (
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
)

type SponsorResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *SponsorResource) List(c buffalo.Context) error {
	sl, err := models.GetSponsorList()
	if err != nil {
		return c.Error(500, err)
	}
	ssl := models.SortedSponsorList(sl)
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("sponsors", ssl)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("sponsors.html"))

}

// Show default implementation.
func (v *SponsorResource) Show(c buffalo.Context) error {
	ids := c.Param("speaker_ID")
	id, err := strconv.Atoi(ids)
	if err != nil {
		return c.Error(400, err)
	}
	p, err := models.GetSponsor(id)
	if err != nil {
		return c.Error(404, err)
	}

	c.Set("sponsor", p)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("sponsor.html"))
}
