package handlers

import (
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}
func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contact")
}
func renderTemplate(w http.ResponseWriter, tpl string) {
	t, err := template.ParseFiles("./templates/" + tpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
