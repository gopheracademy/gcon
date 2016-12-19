package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Location
func Test_Location(t *testing.T) {
	l := &models.Location{}
	l.Name = "Great Divide"
	l.Zip = "85013-9810"
	if m := l.String(); !strings.Contains(m, "Great Divide") {
		t.Errorf("expected contains %s, got %s", "Great Divide", m)
	}
}

// Test_Locations
func Test_Locations(t *testing.T) {
	l := &models.Locations{
		{
			Name: "Great Divide",
			Zip:  "85013-9810",
		},
	}
	if m := l.String(); !strings.Contains(m, "Great Divide") {
		t.Errorf("expected contains %s, got %s", "Great Divide", m)
	}
}
