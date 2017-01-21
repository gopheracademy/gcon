package models_test

import (
	"fmt"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Sponsor
func Test_Sponsor(t *testing.T) {
	sp, err := models.GetSponsor(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sp)
}

// Test_Sponsor
func Test_SponsorAll(t *testing.T) {
	sl, err := models.GetSponsorList()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sl)
}
