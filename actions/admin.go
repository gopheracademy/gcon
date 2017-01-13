package actions

import "github.com/gobuffalo/buffalo"

// AdminPage is a container that holds variables to populate the admin templates
type AdminPage struct {
	Title    string
	Subtitle string
	Page     string
}

// makeAdminPage returns an AdminPage struct struct which populates
// the admin template.  All pages with the admin templates should include
// this by calling c.Set("adminpage", makeAdminPage("Title", "Subtitle", "Page Name"))
func makeAdminPage(title, subtitle, page string) AdminPage {
	return AdminPage{
		Title:    title,
		Subtitle: subtitle,
		Page:     page,
	}
}

// AdminHandler serves the admin root
func AdminHandler(c buffalo.Context) error {
	adminR.HTMLLayout = "main.html"
	c.Set("adminpage", makeAdminPage("Dashboard", "get stuff done!", "Index"))
	return c.Render(200, adminR.HTML("index.html"))
}
