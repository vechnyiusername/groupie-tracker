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

	log.Println("Artists loaded:", len(Artists))
	log.Println("Locations loaded:", len(Locations))
	log.Println("Dates loaded:", len(Dates))
	log.Println("Relations loaded:", len(Relations))
}


func homeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	err := t.Execute(w, Artists)
	if err != nil {
		log.Println("template error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	artist := models.FindArtist(id, Artists)
	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
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

	t := template.Must(template.ParseFiles("templates/artist.html"))
	if err := t.Execute(w, full); err != nil {
		log.Println("template error:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}


func main() {
	// Раздача статики (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Главная страница
	http.HandleFunc("/artist", artistHandler) // сначала специфичный маршрут
	http.HandleFunc("/", homeHandler)         // затем общий

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
