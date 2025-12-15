package models

import (
	"encoding/json"
	"log"
	"os"
)

// -------------------- LOAD ARTISTS --------------------

func LoadArtists() []Artist {
	file, err := os.ReadFile("data/artists.json")
	if err != nil {
		log.Println("Error reading artists.json:", err)
		return nil
	}

	var artists []Artist
	if err := json.Unmarshal(file, &artists); err != nil {
		log.Println("Error unmarshalling artists.json:", err)
		return nil
	}

	return artists
}

// -------------------- LOAD LOCATIONS --------------------

func LoadLocations() []Locations {
	file, err := os.ReadFile("data/locations.json")
	if err != nil {
		log.Println("Error reading locations.json:", err)
		return nil
	}

	var locations []Locations
	if err := json.Unmarshal(file, &locations); err != nil {
		log.Println("Error unmarshalling locations.json:", err)
		return nil
	}

	return locations
}

// -------------------- LOAD DATES --------------------

func LoadDates() []Dates {
	file, err := os.ReadFile("data/dates.json")
	if err != nil {
		log.Println("Error reading dates.json:", err)
		return nil
	}

	var dates []Dates
	if err := json.Unmarshal(file, &dates); err != nil {
		log.Println("Error unmarshalling dates.json:", err)
		return nil
	}

	return dates
}

// -------------------- LOAD RELATIONS --------------------

func LoadRelations() []Relation {
	file, err := os.ReadFile("data/relations.json")
	if err != nil {
		log.Println("Error reading relations.json:", err)
		return nil
	}

	var relations []Relation
	if err := json.Unmarshal(file, &relations); err != nil {
		log.Println("Error unmarshalling relations.json:", err)
		return nil
	}

	return relations
}
