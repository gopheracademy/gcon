package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

func Test_Speaker(t *testing.T) {
	s := &models.Speaker{}
	s.FirstName = "Gary"
	s.LastName = "Gopher"
	if m := s.String(); !strings.Contains(m, "Gary") {
		t.Errorf("expected contains %s, got %s", "Gary", m)
	}
}
