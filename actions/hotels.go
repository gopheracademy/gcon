package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
)

func HotelsHandler(c buffalo.Context) error {
	hp, err := models.GetHotelList()
	if err != nil {
		return c.Error(500, err)
	}
	c.Set("hotels", hp)
	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("hotels.html"))
}
