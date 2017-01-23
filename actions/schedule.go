package actions

import "github.com/gobuffalo/buffalo"

type ScheduleResource struct {
	buffalo.Resource
}

// List default implementation.
func (v *ScheduleResource) List(c buffalo.Context) error {
	return c.Render(200, publicR.HTML("soon.html"))

}

// Show default implementation.
func (v *ScheduleResource) Show(c buffalo.Context) error {
	return c.Render(200, publicR.String("Schedule#Show"))
}
