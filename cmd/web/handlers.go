package main

import (
	"log"
	"net/http"
	"text/template"

	"aliyeysides/personal-site/ui"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts := template.Must(template.ParseFS(ui.Templates, "html/*"))

	err := ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
