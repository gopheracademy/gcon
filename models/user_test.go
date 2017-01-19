package models_test

import (
	"strings"
	"testing"

	"github.com/gopheracademy/gcon/models"
)

func Test_User(t *testing.T) {
	u := &models.User{}
	u.FirstName = "Brian"
	u.Password = "password"
	if m := u.String(); !strings.Contains(m, "Brian") {
		t.Errorf("expected contains %s, got %s", "Brian", m)
	}
}
