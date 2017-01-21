package models

import "github.com/gopheracademy/gccms/content"

type SponsorList struct {
	Diamond  []content.Sponsor
	Platinum []content.Sponsor
	Gold     []content.Sponsor
	Silver   []content.Sponsor
	Bronze   []content.Sponsor
}

func SortedSponsorList(slr []content.Sponsor) SponsorList {
	var sl SponsorList
	for _, s := range slr {
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
	return sl

}
