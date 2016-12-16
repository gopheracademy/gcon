package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_PresentationSpeaker
func Test_PresentationSpeaker(t *testing.T) {
	p := &models.PresentationSpeaker{}
	p.PresentationID = 1
	p.SpeakerID = 1
	if m := p.String(); !strings.Contains(m, "1") {
		t.Errorf("expected contains %s, got %s", "1", m)
	}
}
