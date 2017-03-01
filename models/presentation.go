package models

import (
	"fmt"

	"github.com/gopheracademy/gccms/content"
)

type Presentation struct {
	Presentation content.Presentation
	Speakers     []content.Speaker
}

// GetFullPresentation returns a full presentation, with
// the speaker populated.
func GetFullPresentation(id int) (Presentation, error) {
	var p Presentation
	pl, err := GetPresentation(id)
	if err != nil {
		return p, err
	}
	p.Presentation = pl
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

func GetPresentations() []Presentation {
	var pp []Presentation
	pl, err := GetPresentationList()
	if err != nil {
		fmt.Println(err)
		return pp
	}

	for _, p := range pl {
		var pr Presentation
		pr.Presentation = p
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
