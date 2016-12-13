package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Hotel
func Test_Hotel(t *testing.T) {
	h := &models.Hotel{}
	h.BlockCode = "GOPHERCON2017"
	h.RoomRate = 225.0
	if m := h.String(); !strings.Contains(m, "GOPHERCON2017") {
		t.Errorf("expected contains %s, got %s", "GOPHERCON2017", m)
	}
}
