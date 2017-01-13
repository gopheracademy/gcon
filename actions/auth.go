package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gopheracademy/gcon/models"
	"golang.org/x/crypto/bcrypt"
)

// AdminHandler serves the admin root
func LoginHandler(c buffalo.Context) error {
	adminR.HTMLLayout = "login.html"
	return c.Render(200, adminR.HTML("index.html"))
}

func AuthHandler(c buffalo.Context) error {

	// auth request
	username := c.Request().FormValue("username")
	pass := c.Request().FormValue("password")
	var u models.User
	err := models.DB.Where("email = ?", username).First(&u)
	if err != nil {

		adminR.HTMLLayout = "login.html"
		c.Set("HasLoginError", true)
		c.Set("LoginError", "Invalid Login")
		return c.Render(401, adminR.HTML("login.html"))
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pass)); err != nil {

		adminR.HTMLLayout = "login.html"
		c.Set("HasLoginError", true)
		c.Set("LoginError", "Invalid Login")
		return c.Render(401, adminR.HTML("login.html"))
	}

	// if OK, set username in session
	c.Session().Set("username", c.Request().FormValue("username"))
	c.Session().Save()

	adminR.HTMLLayout = "main.html"
	c.Set("user", u)
	c.Set("adminpage", makeAdminPage("Dashboard", "get stuff done!", "Index"))
	return c.Render(200, adminR.HTML("index.html"))
}
