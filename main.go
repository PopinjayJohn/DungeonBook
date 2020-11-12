package main

import (
	"github.com/PopinjayJohn/DungeonBook/views"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var index *views.View

type PageData struct {
  PageTitle string
}

func main() {
	staticDir := "static"

	// Create a router
	r := mux.NewRouter()

	// Set up static files and basic resources TODO; This could be nicely packed up in a custom handler
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	r.HandleFunc("/favicon.ico", FaviconHandler)
	r.HandleFunc("/icon.png", IconHandler)
	r.HandleFunc("/site.webmanifest", WebmanifestHandler)

	// Set up routes
	index = views.NewView("base", "views/index.gohtml")
	r.HandleFunc("/", HomeHandler)

	// Start serving
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
}
func IconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/icon.png")
}
func WebmanifestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/site.webmanifest")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  data := PageData{
    PageTitle: "DungeonBook - For the social player",
  }
	index.Render(w, data)
}
