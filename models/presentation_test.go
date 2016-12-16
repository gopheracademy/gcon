package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Presentation
func Test_Presentation(t *testing.T) {
	p := &models.Presentation{}
	p.Title = "Go-ing The Distance"
	p.ShortDescription = "Building Charlie's Go Blog"
	if m := p.String(); !strings.Contains(m, "Distance") {
		t.Errorf("expected contains %s, got %s", "Distance", m)
	}
}
