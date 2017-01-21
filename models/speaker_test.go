package models_test

import (
	"fmt"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Speaker
func Test_Speaker(t *testing.T) {
	sp, err := models.GetSpeaker(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sp)
}

// Test_Speaker
func Test_SpeakerAll(t *testing.T) {
	sp, err := models.GetSpeakerList()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sp)
}
