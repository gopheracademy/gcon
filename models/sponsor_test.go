package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Sponsor
func Test_Sponsor(t *testing.T) {
	s := &models.Sponsor{}
	s.Name = "Google"
	s.Website = "www.golang.org"
	if m := s.String(); !strings.Contains(m, "Google") {
		t.Errorf("expected contains %s, got %s", "Google", m)
	}
}

// Test_Sponsors
func Test_Sponsors(t *testing.T) {
	s := &models.Sponsors{
		{
			Name:    "Google",
			Website: "www.golang.org",
		},
	}
	if m := s.String(); !strings.Contains(m, "Google") {
		t.Errorf("expected contains %s, got %s", "Google", m)
	}
}
