package actions_test

import (
	"testing"

	"github.com/gopheracademy/gcon/actions"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

func Test_HomeHandler(t *testing.T) {
	r := require.New(t)

	w := willie.New(actions.App())
	res := w.Request("/").Get()

	r.Equal(200, res.Code)
	r.Contains(res.Body.String(), "Largest event in the world dedicated")
	r.Contains(res.Body.String(), "Interested in speaking?")
}
