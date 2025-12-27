package main

import (
	"encoding/json"
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var (
	Artists   []models.Artist
	Locations []models.LocationItem
	Dates     []models.DateItem
	Relations []models.RelationItem

	dataMu sync.RWMutex
)

func init() {
	// load data from API (with cache fallback). functions are resilient and won't panic.
	Artists = models.LoadArtists()
	Locations = models.LoadLocations()
	Dates = models.LoadDates()
	Relations = models.LoadRelations()
}

type ErrorPageData struct {
	StatusCode int
	Title      string
	Message    string
}

func renderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	data := ErrorPageData{
		StatusCode: status,
	}

	switch status {
	case http.StatusNotFound:
		data.Title = "404 Not Found"
		data.Message = "The page is not found or the resource is missing."
	case http.StatusInternalServerError:
		data.Title = "500 Internal Server Error"
		data.Message = "Something went wrong on the server side. Try again later."
	default:
		data.Title = "Error"
		data.Message = "Error occured."
	}

	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		log.Println("Error parsing error template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, data); err != nil {
		log.Println("Error executing error template:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles(
		"templates/layout.html",
		"templates/index.html",
	)
	if err != nil {
		log.Println("Error parsing home templates:", err)
		renderError(w, http.StatusInternalServerError)
		return
	}

	// safely read artists
	dataMu.RLock()
	artists := Artists
	dataMu.RUnlock()

	err = tmpl.ExecuteTemplate(w, "layout", artists)
	if err != nil {
		log.Println("template error:", err)
		renderError(w, http.StatusInternalServerError)
		return
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		renderError(w, http.StatusNotFound)
		return
	}

	dataMu.RLock()
	artists := Artists
	locs := Locations
	dts := Dates
	rels := Relations
	dataMu.RUnlock()

	artist := models.FindArtist(id, artists)
	if artist == nil {
		renderError(w, http.StatusNotFound)
		return
	}

	locations := models.FindLocations(id, locs)
	dates := models.FindDates(id, dts)
	relations := models.FindRelation(id, rels)

	full := models.ArtistFull{
		Artist:         *artist,
		Locations:      locations,
		Dates:          dates,
		DatesLocations: relations,
	}

	tmpl, err := template.ParseFiles("templates/layout.html", "templates/artist.html")
	if err != nil {
		log.Println("Error parsing artist template:", err)
		renderError(w, http.StatusInternalServerError)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "layout", full); err != nil {
		log.Println("Template execution error:", err)
		renderError(w, http.StatusInternalServerError)
	}
}

func concertsAPIHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	dataMu.RLock()
	rels := Relations
	dataMu.RUnlock()

	res := models.FindRelation(id, rels)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("Error encoding concerts json:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/artist", artistHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	// API used by frontend (no direct external API calls from browser)
	http.HandleFunc("/api/concerts", concertsAPIHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
