package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

// Test_Contact
func Test_Contact(t *testing.T) {
	c := &models.Contact{}
	c.Name = "Hyatt"
	c.Phone = "(303) 436-1234"
	if m := c.String(); !strings.Contains(m, "Hyatt") {
		t.Errorf("expected contains %s, got %s", "Hyatt", m)
	}
}
