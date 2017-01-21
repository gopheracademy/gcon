package models

import (
	"errors"
	"time"

	"github.com/bketelsen/ponzi"
	"github.com/gopheracademy/gccms/content"
)

var sponsorCache *ponzi.Cache

func initSponsorCache() {
	if sponsorCache == nil {
		sponsorCache = ponzi.New("http://127.0.0.1:8080", 1*time.Minute, 30*time.Second)
	}
}

type SponsorList struct {
	Diamond  []content.Sponsor
	Platinum []content.Sponsor
	Gold     []content.Sponsor
	Silver   []content.Sponsor
	Bronze   []content.Sponsor
}

func GetSponsor(id int) (content.Sponsor, error) {
	initSponsorCache()
	var sp content.SponsorListResult
	err := sponsorCache.Get(id, "Sponsor", &sp)
	if err != nil {
		return content.Sponsor{}, err
	}
	if len(sp.Data) == 0 {
		return content.Sponsor{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetSponsorList() (SponsorList, error) {
	initSponsorCache()
	var sp content.SponsorListResult
	err := sponsorCache.GetAll("Sponsor", &sp)
	if err != nil {
		return SponsorList{}, err
	}
	if len(sp.Data) == 0 {
		return SponsorList{}, errors.New("Not Found")
	}
	var sl SponsorList
	for _, s := range sp.Data {
		switch s.Level {
		case "Diamond":
			sl.Diamond = append(sl.Diamond, s)
		case "Platinum":
			sl.Platinum = append(sl.Platinum, s)
		case "Gold":
			sl.Gold = append(sl.Gold, s)
		case "Silver":
			sl.Silver = append(sl.Silver, s)
		case "Bronze":
			sl.Bronze = append(sl.Bronze, s)
		}
	}
	return sl, err

}
