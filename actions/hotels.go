package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
)

type HotelsResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *HotelsResource) List(c buffalo.Context) error {

	var ii models.Hotels
	models.DB.All(&ii)

	for x, i := range ii {
		var l models.Location
		err := models.DB.Find(&l, i.LocationID)
		if err != nil {
			return c.Error(500, err)
		}
		i.Location = &l
		ii[x] = i
	}
	c.Set("hotels", ii)
	return c.Render(200, publicR.HTML("hotels.html"))

}

// Show default implementation.
func (v *HotelsResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Hotels#Show"))
}
