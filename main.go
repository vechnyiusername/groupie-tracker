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

type ErrorPageData struct {
	StatusCode int
	Title      string
	Message    string
}

// универсальная функция для отображения ошибок
func renderError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	data := ErrorPageData{
		StatusCode: status,
	}

	switch status {
	case http.StatusNotFound:
		data.Title = "404 Not Found"
		data.Message = "Страница не найдена или ресурс отсутствует."
	case http.StatusInternalServerError:
		data.Title = "500 Internal Server Error"
		data.Message = "На сервере что-то пошло не так. Попробуйте позже."
	default:
		data.Title = "Error"
		data.Message = "Произошла ошибка."
	}

	t := template.Must(template.ParseFiles("templates/error.html"))
	if err := t.Execute(w, data); err != nil {
		// если даже шаблон ошибки сломался — отправим простой текст
		http.Error(w, "critical error", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound)
		return
	}

	log.Println("Home handler, artists count:", len(Artists))

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

	tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/artist.html"))

	if err := tmpl.ExecuteTemplate(w, "layout", full); err != nil {
		log.Println("template error:", err)
		renderError(w, http.StatusInternalServerError)
		return
	}
}

func main() {
	// Раздача статики (CSS, JS)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Главная страница
	http.HandleFunc("/artist", artistHandler)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
