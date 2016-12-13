package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Hotel
func Test_Hotel(t *testing.T) {
	h := &models.Hotel{}
	h.PropertyName = "Hilton"
	h.Address = "123 State Street"
	if m := h.String(); !strings.Contains(m, "Hilton") {
		t.Errorf("expected contains %s, got %s", "Hilton", m)
	}
}
