package models

import (
	"encoding/json"
	"log"
	"os"
)

// ---------- ARTISTS ----------

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

// ---------- LOCATIONS ----------

func LoadLocations() []LocationItem {
	file, err := os.ReadFile("data/locations.json")
	if err != nil {
		log.Println("Error reading locations.json:", err)
		return nil
	}

	var container LocationsIndex
	if err := json.Unmarshal(file, &container); err != nil {
		log.Println("Error unmarshalling locations.json:", err)
		return nil
	}

	return container.Index
}

// ---------- DATES ----------

func LoadDates() []DateItem {
	file, err := os.ReadFile("data/dates.json")
	if err != nil {
		log.Println("Error reading dates.json:", err)
		return nil
	}

	var container DatesIndex
	if err := json.Unmarshal(file, &container); err != nil {
		log.Println("Error unmarshalling dates.json:", err)
		return nil
	}

	return container.Index
}

// ---------- RELATIONS ----------

func LoadRelations() []RelationItem {
	file, err := os.ReadFile("data/relations.json")
	if err != nil {
		log.Println("Error reading relations.json:", err)
		return nil
	}

	var container RelationsIndex
	if err := json.Unmarshal(file, &container); err != nil {
		log.Println("Error unmarshalling relations.json:", err)
		return nil
	}

	return container.Index
}
