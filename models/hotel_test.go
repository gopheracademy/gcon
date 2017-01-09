package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Hotels
func Test_Hotels(t *testing.T) {
	h := &models.Hotels{
		{
			BlockCode: "GOPHERCON2017",
			RoomRate:  225,
		},
	}
	if m := h.String(); !strings.Contains(m, "GOPHERCON2017") {
		t.Errorf("expected contains %s, got %s", "GOPHERCON2017", m)
	}
}
