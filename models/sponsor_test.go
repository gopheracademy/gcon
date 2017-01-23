package models_test

import (
	"fmt"
	"testing"

	"github.com/gopheracademy/gcon/models"
	"github.com/kr/pretty"
)

// Test_Sponsor
func Test_Sponsor(t *testing.T) {
	models.BaseURL = "http://127.0.0.1:8080"
	sp, err := models.GetSponsor(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sp)
}

// Test_Sponsor
func Test_SponsorAll(t *testing.T) {

	models.BaseURL = "http://127.0.0.1:8080"
	sl, err := models.GetSponsorList()
	if err != nil {
		t.Error(err)
	}
	ssl := models.SortedSponsorList(sl)
	pretty.Println(ssl)
}
