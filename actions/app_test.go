package actions

import (
	"os"
	"testing"
)

// TestInit
func TestInit(t *testing.T) {
	os.Setenv("GO_ENV", "development")
	ENV = os.Getenv("GO_ENV")
}

// TestApp
func TestApp(t *testing.T) {

}

// TestSetTemplate
func TestSetTemplate(t *testing.T) {

}
