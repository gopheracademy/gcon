package main

import (
	"log"
	"net/http"

	"github.com/gopheracademy/gcon/actions"
)

func main() {
	log.Fatal(http.ListenAndServe(":3000", actions.App()))
}
