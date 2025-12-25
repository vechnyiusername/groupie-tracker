package main

import (
	"groupie-tracker/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	Artists   []models.Artist
	Locations []models.LocationItem
	Dates     []models.DateItem
	Relations []models.RelationItem
)

func init() {
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

	t := template.Must(template.ParseFiles("templates/error.html"))
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "critical error", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/index.html"))

	err := tmpl.ExecuteTemplate(w, "layout", Artists)
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

	artist := models.FindArtist(id, Artists)
	if artist == nil {
		renderError(w, http.StatusNotFound)
		return
	}

	locations := models.FindLocations(id, Locations)
	dates := models.FindDates(id, Dates)
	relations := models.FindRelation(id, Relations)

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

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/artist", artistHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
