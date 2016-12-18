package actions_test

import (
	"testing"

	"github.com/gopheracademy/gcon/actions"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

// TestMakeAdminPage
/*func TestMakeAdminPage(t *testing.T) {
	ap := actions.makeAdminPage("Admin Page", "Where you do admin stuff...", "page something")
	b, _ := json.Marshal(ap)
	if !strings.Contains(string(b), "Admin Page") {
		t.Errorf("expected contains %s, got %s", "Admin Page", string(b))
	}
}*/

// Test_AdminHandler
func Test_AdminHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(actions.App())
	res := w.Request("/").Get()

	r.Equal(200, res.Code)
	r.Contains(res.Body.String(), "GopherCon")
}
