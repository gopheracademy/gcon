package models

import (
	"fmt"
	"sort"

	"github.com/gopheracademy/gccms/content"
)

type Presentation struct {
	Presentation content.Presentation
	Speakers     []content.Speaker
	Slot         content.Slot
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
	sid, err := getID(p.Presentation.Slot)
	slot, err := GetSlot(sid)
	p.Slot = slot

	return p, nil
}

func GetPresentations() []Presentation {
	var pp []Presentation
	pl, err := GetPresentationList()
	if err != nil {
		fmt.Println(err)
		return pp
	}

	sort.Slice(pl, func(i, j int) bool { return pl[i].DisplayOrder < pl[j].DisplayOrder })
	for _, p := range pl {
		var pr Presentation
		pr.Presentation = p
		sid, err := getID(pr.Presentation.Slot)
		slot, err := GetSlot(sid)

		if err != nil {
			fmt.Println(err, sid, slot)
		}
		pr.Slot = slot
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
