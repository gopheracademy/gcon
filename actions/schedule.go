package actions

import (
	"sort"

	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
)

func ScheduleHandler(c buffalo.Context) error {

	presentations := models.GetPresentations()
	c.Set("presentations", presentations)

	workshops := models.GetWorkshops()
	c.Set("workshops", workshops)

	dayone, daytwo := byDays(presentations)

	sort.Slice(dayone, func(i, j int) bool { return dayone[i].Slot.Number < dayone[j].Slot.Number })
	sort.Slice(daytwo, func(i, j int) bool { return daytwo[i].Slot.Number < daytwo[j].Slot.Number })

	c.Set("dayone", dayone)
	c.Set("daytwo", daytwo)

	publicR.HTMLLayout = "main.html"
	return c.Render(200, publicR.HTML("schedule.html"))
}

func byDays(p []models.Presentation) ([]models.Presentation, []models.Presentation) {

	var day1 []models.Presentation
	var day2 []models.Presentation

	for _, pres := range p {
		if pres.Presentation.Day == 1 {
			day1 = append(day1, pres)
		}

		if pres.Presentation.Day == 2 {
			day2 = append(day2, pres)
		}
	}

	return day1, day2

}
