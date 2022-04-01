package main

import (
	"net/http"
	"text/template"
)

type Chatrapti struct {
	Name string
	Age  int
}

func main() {
	raje := Chatrapti{"raje", 12}
	template := template.Must(template.ParseFiles("template/template.html"))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			raje.Name = name
		}
		if err := template.ExecuteTemplate(w, "template.html", raje); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
