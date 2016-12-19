package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Speaker
func Test_Speaker(t *testing.T) {
	s := &models.Speaker{}
	s.FirstName = "Brian"
	s.PhotoURL = "https://blog.gopheracademy.com/postimages/advent-2015/gophercon2016.svg"
	if m := s.String(); !strings.Contains(m, "Brian") {
		t.Errorf("expected contains %s, got %s", "Brian", m)
	}
}

// Test_Speakers
func Test_Speakers(t *testing.T) {
	s := &models.Speakers{
		{
			FirstName: "Brian",
			PhotoURL:  "https://blog.gopheracademy.com/postimages/advent-2015/gophercon2016.svg",
		},
	}
	if m := s.String(); !strings.Contains(m, "Brian") {
		t.Errorf("expected contains %s, got %s", "Brian", m)
	}
}
