/*
* CODE GENERATED AUTOMATICALLY WITH github.com/bketelsen/ponzigen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package models

import (
	"errors"
	"time"

	"github.com/bketelsen/ponzi"
	"github.com/gopheracademy/gccms/content"
)

var BaseURL string

var eventCache *ponzi.Cache
var hotelCache *ponzi.Cache
var pageCache *ponzi.Cache
var presentationCache *ponzi.Cache
var speakerCache *ponzi.Cache
var sponsorCache *ponzi.Cache

func initEventCache() {
	if eventCache == nil {
		eventCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initHotelCache() {
	if hotelCache == nil {
		hotelCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initPageCache() {
	if pageCache == nil {
		pageCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initPresentationCache() {
	if presentationCache == nil {
		presentationCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initSpeakerCache() {
	if speakerCache == nil {
		speakerCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initSponsorCache() {
	if sponsorCache == nil {
		sponsorCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}

func GetEvent(id int) (content.Event, error) {
	initEventCache()
	var sp content.EventListResult
	err := eventCache.Get(id, "Event", &sp)
	if err != nil {
		return content.Event{}, err
	}
	if len(sp.Data) == 0 {
		return content.Event{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetHotel(id int) (content.Hotel, error) {
	initHotelCache()
	var sp content.HotelListResult
	err := hotelCache.Get(id, "Hotel", &sp)
	if err != nil {
		return content.Hotel{}, err
	}
	if len(sp.Data) == 0 {
		return content.Hotel{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetPage(id int) (content.Page, error) {
	initPageCache()
	var sp content.PageListResult
	err := pageCache.Get(id, "Page", &sp)
	if err != nil {
		return content.Page{}, err
	}
	if len(sp.Data) == 0 {
		return content.Page{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetPresentation(id int) (content.Presentation, error) {
	initPresentationCache()
	var sp content.PresentationListResult
	err := presentationCache.Get(id, "Presentation", &sp)
	if err != nil {
		return content.Presentation{}, err
	}
	if len(sp.Data) == 0 {
		return content.Presentation{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetSpeaker(id int) (content.Speaker, error) {
	initSpeakerCache()
	var sp content.SpeakerListResult
	err := speakerCache.Get(id, "Speaker", &sp)
	if err != nil {
		return content.Speaker{}, err
	}
	if len(sp.Data) == 0 {
		return content.Speaker{}, errors.New("Not Found")
	}
	return sp.Data[0], err

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

func GetEventList() ([]content.Event, error) {
	initEventCache()
	var sp content.EventListResult
	err := eventCache.GetAll("Event", &sp)
	if err != nil {
		return []content.Event{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Event{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetHotelList() ([]content.Hotel, error) {
	initHotelCache()
	var sp content.HotelListResult
	err := hotelCache.GetAll("Hotel", &sp)
	if err != nil {
		return []content.Hotel{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Hotel{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetPageList() ([]content.Page, error) {
	initPageCache()
	var sp content.PageListResult
	err := pageCache.GetAll("Page", &sp)
	if err != nil {
		return []content.Page{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Page{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetPresentationList() ([]content.Presentation, error) {
	initPresentationCache()
	var sp content.PresentationListResult
	err := presentationCache.GetAll("Presentation", &sp)
	if err != nil {
		return []content.Presentation{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Presentation{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetSpeakerList() ([]content.Speaker, error) {
	initSpeakerCache()
	var sp content.SpeakerListResult
	err := speakerCache.GetAll("Speaker", &sp)
	if err != nil {
		return []content.Speaker{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Speaker{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetSponsorList() ([]content.Sponsor, error) {
	initSponsorCache()
	var sp content.SponsorListResult
	err := sponsorCache.GetAll("Sponsor", &sp)
	if err != nil {
		return []content.Sponsor{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Sponsor{}, errors.New("Not Found")
	}
	return sp.Data, err

}
