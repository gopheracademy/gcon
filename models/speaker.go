package models

import (
	"errors"
	"time"

	"github.com/bketelsen/ponzi"
	"github.com/gopheracademy/gccms/content"
)

var speakerCache *ponzi.Cache

func initSpeakerCache() {
	if speakerCache == nil {
		speakerCache = ponzi.New("http://127.0.0.1:8080", 1*time.Minute, 30*time.Second)
	}
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
