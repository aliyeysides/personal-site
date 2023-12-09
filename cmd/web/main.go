package main

import (
	"fmt"
	"log"
	"net/http"

	"aliyeysides/personal-site/templ"
	"aliyeysides/personal-site/ui"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(ui.Resources))
	mux.Handle("/static/", fileServer)

	appComponent := templates.HomePage()
	aboutComponent := templates.AboutPage()
	myWorkComponent := templates.MyWorkPage()

	mux.Handle("/", templ.Handler(appComponent))
	mux.Handle("/about-site", templ.Handler(aboutComponent))
	mux.Handle("/my-work", templ.Handler(myWorkComponent))

	fmt.Println("Listening on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
