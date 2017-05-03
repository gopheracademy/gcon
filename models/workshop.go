package models

import (
	"fmt"

	"github.com/gopheracademy/gccms/content"
)

type Workshop struct {
	Workshop content.Workshop
	Speakers []content.Speaker
}

// GetFullWorkshop returns a full presentation, with
// the speaker populated.
func GetFullWorkshop(id int) (Workshop, error) {
	var p Workshop
	pl, err := GetWorkshop(id)
	if err != nil {
		return p, err
	}
	p.Workshop = pl
	for _, s := range pl.Speakers {
		id, err := getID(s)
		if err != nil {
			fmt.Println(err, id, s)
			continue
		}
		sp, err := GetSpeaker(id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p.Speakers = append(p.Speakers, sp)
	}

	return p, nil
}

func GetWorkshops() []Workshop {
	var pp []Workshop
	pl, err := GetWorkshopList()
	if err != nil {
		fmt.Println(err)
		return pp
	}

	for _, p := range pl {
		var pr Workshop
		pr.Workshop = p
		for _, s := range p.Speakers {
			id, err := getID(s)
			if err != nil {
				fmt.Println(err, id, s)
				continue
			}
			sp, err := GetSpeaker(id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			pr.Speakers = append(pr.Speakers, sp)
		}
		pp = append(pp, pr)
	}
	return pp
}
