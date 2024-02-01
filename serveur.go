package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "Custom 404")
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func handleFunctions(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		tpl.ExecuteTemplate(w, "index.html", nil)
	case "/Reglages":
		tpl.ExecuteTemplate(w, "Reglages.html", nil)
	case "/gamepage":
		tpl.ExecuteTemplate(w, "gamepage.html", nil)
	case "/test":
		tpl.ExecuteTemplate(w, "test.html", nil)
	default:
		errorHandler(w, r, http.StatusNotFound)
	}
}

func main() {
	// Gestionnaire pour la page principale
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", handleFunctions)
	http.HandleFunc("/Reglages", handleFunctions)
	http.HandleFunc("/gamepage", handleFunctions)
	http.HandleFunc("/test", handleFunctions)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
