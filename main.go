package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gopheracademy/gcon/actions"
	"github.com/gopheracademy/gcon/models"
	"github.com/markbates/going/defaults"
)

func main() {

	port := defaults.String(os.Getenv("PORT"), "3002")
	baseURL := defaults.String(os.Getenv("CMS_URL"), "https://cms.gophercon.com")
	models.BaseURL = baseURL
	log.Printf("Starting gcon on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), actions.App()))
}
