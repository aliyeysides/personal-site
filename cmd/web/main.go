package main

import (
	"fmt"
	"log"
	"net/http"

	"aliyeysides/personal-site/ui"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(ui.Resources))
	mux.Handle("/static/", fileServer)

	mux.HandleFunc("/", home)

	fmt.Println("Listening on port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
