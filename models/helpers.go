package models

import (
	"net/url"
	"strconv"

	"github.com/gopheracademy/gccms/content"
)

type HomePage struct {
	Hotels        []content.Hotel
	Sponsors      SponsorList
	Presentations []Presentation
}

func GetHomePage() (*HomePage, error) {
	var hp HomePage
	hotels, err := GetHotelList()
	if err != nil {
		return &hp, err
	}
	hp.Hotels = hotels

	presentations := GetPresentations()

	hp.Presentations = presentations
	sponsors, err := GetSponsorList()
	if err != nil {
		return &hp, err
	}
	hp.Sponsors = SortedSponsorList(sponsors)
	return &hp, nil
}

func getID(s string) (int, error) {
	//?type=Module&id=4
	u, err := url.Parse(s)
	if err != nil {
		return 0, err
	}
	vals := u.Query()
	ii, ok := vals["id"]
	if !ok {
		return 0, err
	}
	return strconv.Atoi(ii[0])
}
