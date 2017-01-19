package grifts

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"

	"github.com/gopheracademy/gcon/models"
	"github.com/kr/pretty"
	. "github.com/markbates/grift/grift"
)

var _ = Set("admin", func(c *Context) error {
	if len(os.Args) < 6 {
		return fmt.Errorf("Usage: admin [email] [password] [first_name] [last_name]")
	}
	email := os.Args[2]
	pass := os.Args[3]
	first := os.Args[4]
	last := os.Args[5]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Error hashing password: %s", err)
	}
	fmt.Println(string(hashedPassword))
	u := models.User{
		FirstName: first,
		LastName:  last,
		Email:     email,
		Password:  string(hashedPassword),
		Role:      "Admin",
	}
	err = models.DB.Save(&u)
	models.DB.Reload(&u)
	fmt.Printf("Created User %d - %s\n", u.ID, u.Email)
	return nil
})

var _ = Set("speaker", func(c *Context) error {
	if len(os.Args) < 8 {
		return fmt.Errorf("Usage: speaker [first_name] [last_name] [bio] [website] [twitter] [linkedin] [facebook] [photo_url]")
	}
	first := os.Args[2]
	last := os.Args[3]
	bio := os.Args[4]
	website := os.Args[5]
	twitter := os.Args[6]
	linkedin := os.Args[7]
	facebook := os.Args[8]
	photoURL := os.Args[9]

	s := models.Speaker{
		FirstName: first,
		LastName:  last,
		Bio:       bio,
		Website:   website,
		Twitter:   twitter,
		Linkedin:  linkedin,
		Facebook:  facebook,
		PhotoURL:  photoURL,
	}
	err = models.DB.Save(&s)
	models.DB.Reload(&s)
	fmt.Printf("Created Speaker %d - %s %s\n", s.ID, s.FirstName, s.LastName)
	return nil
})

var _ = Set("hotels", func(c *Context) error {
	var ii models.Hotels
	models.DB.All(&ii)

	for x, i := range ii {
		var l models.Location
		err := models.DB.Find(&l, i.LocationID)
		if err != nil {
			fmt.Println(err)
			continue
		}
		i.Location = &l
		ii[x] = i
	}
	pretty.Println(ii)
	return nil
})

var _ = Set("locations", func(c *Context) error {
	var ii models.Locations
	models.DB.All(&ii)

	fmt.Println(ii)
	return nil
})
var _ = Add("seed", func(c *Context) error {

	hgiloc := models.Location{}
	hgiloc.Address = "1400 Welton Street"
	hgiloc.City = "Denver"
	hgiloc.State = "CO"
	hgiloc.Zip = "80202"

	err := models.DB.Save(&hgiloc)
	if err != nil {
		return err
	}
	models.DB.Reload(&hgiloc)

	hgi := models.Hotel{}
	hgi.Name = "Hilton Garden Inn"
	hgi.Phone = "855.215.1283 (refer to GopherCon)"
	hgi.Description = "The beautiful Hilton Garden Inn will be a comfortable place to stay for you."
	hgi.BlockCode = "http://hiltongardeninn.hilton.com/en/gi/groups/personalized/D/DENDDGI-GOPHER-20170711/index.jhtml?WT.mc_id=POG"
	hgi.PhotoURL = "/assets/public/img/hotels/hotel-hgi.jpg"
	hgi.Onsale = true
	hgi.RoomRate = 195
	hgi.LocationID = hgiloc.ID

	err = models.DB.Save(&hgi)
	if err != nil {
		return err
	}
	models.DB.Reload(&hgi)

	hploc := models.Location{}
	hploc.Address = "440 14th Street"
	hploc.City = "Denver"
	hploc.State = "CO"
	hploc.Zip = "80202"

	err = models.DB.Save(&hploc)
	if err != nil {
		return err
	}
	models.DB.Reload(&hploc)

	hp := models.Hotel{}
	hp.Name = "Hyatt Place Denver Downtown"
	hp.Phone = "888.591.1234 (refer to GopherCon)"
	hp.BlockCode = "https://aws.passkey.com/go/GopherAcademyHyattPlace2017"
	hp.PhotoURL = "/assets/public/img/hotels/hyatt.jpg"
	hp.Onsale = true
	hp.RoomRate = 189
	hp.LocationID = hploc.ID

	err = models.DB.Save(&hp)
	if err != nil {
		return err
	}
	models.DB.Reload(&hp)

	hhloc := models.Location{}
	hhloc.Address = "440 14th Street"
	hhloc.City = "Denver"
	hhloc.State = "CO"
	hhloc.Zip = "80202"

	err = models.DB.Save(&hhloc)
	if err != nil {
		return err
	}
	models.DB.Reload(&hhloc)

	hh := models.Hotel{}
	hh.Name = "Hyatt House Denver Downtown"
	hh.Phone = "888.591.1234 (refer to GopherCon)"
	hh.BlockCode = "https://aws.passkey.com/go/GopherAcademyHyattHouse2017"
	hh.PhotoURL = "/assets/public/img/hotels/hyatt.jpg"
	hh.Onsale = true
	hh.RoomRate = 189
	hh.LocationID = hhloc.ID

	err = models.DB.Save(&hh)
	if err != nil {
		return err
	}
	models.DB.Reload(&hh)

	alloc := models.Location{}
	alloc.Address = "800 15th Street"
	alloc.City = "Denver"
	alloc.State = "CO"
	alloc.Zip = "80202"

	err = models.DB.Save(&alloc)
	if err != nil {
		return err
	}
	models.DB.Reload(&alloc)

	al := models.Hotel{}
	al.Name = "ALoft Downtown Denver"
	al.Phone = "877.462.5638 (refer to GopherCon)"
	al.BlockCode = "https://www.starwoodmeeting.com/events/start.action?id=1612127103&key=33F262BB"
	al.PhotoURL = "/assets/public/img/hotels/aloft.jpg"
	al.Onsale = true
	al.RoomRate = 199
	al.LocationID = alloc.ID

	err = models.DB.Save(&al)
	if err != nil {
		return err
	}
	models.DB.Reload(&al)
	return nil

})
