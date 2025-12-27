package models

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const apiBase = "https://groupietrackers.herokuapp.com/api"

var httpClient = &http.Client{
	Timeout: 6 * time.Second,
}

func fetchAndCache(url, cachePath string, target interface{}) error {
	// try fetching from remote API
	resp, err := httpClient.Get(url)
	if err == nil && resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			body, err := io.ReadAll(resp.Body)
			if err == nil {
				// try unmarshal
				if err := json.Unmarshal(body, target); err == nil {
					// write cache file (best-effort)
					_ = os.WriteFile(cachePath, body, 0644)
					return nil
				}
				log.Println("Error unmarshalling remote response from", url, ":", err)
			} else {
				log.Println("Error reading response body from", url, ":", err)
			}
		} else {
			log.Printf("Non-200 from %s: %d\n", url, resp.StatusCode)
		}
	} else {
		log.Println("Error fetching", url, ":", err)
	}

	// fallback: try reading cached file
	data, err := os.ReadFile(cachePath)
	if err != nil {
		log.Println("No cache available for", cachePath, ":", err)
		return err
	}
	if err := json.Unmarshal(data, target); err != nil {
		log.Println("Error unmarshalling cached file", cachePath, ":", err)
		return err
	}
	log.Println("Loaded data from cache:", cachePath)
	return nil
}

func LoadArtists() []Artist {
	url := apiBase + "/artists"
	var artists []Artist
	if err := fetchAndCache(url, "data/artists.json", &artists); err != nil {
		log.Println("LoadArtists: returning empty artists due to error:", err)
		return nil
	}
	return artists
}

func LoadLocations() []LocationItem {
	url := apiBase + "/locations"
	var container LocationsIndex
	if err := fetchAndCache(url, "data/locations.json", &container); err != nil {
		log.Println("LoadLocations: returning empty locations due to error:", err)
		return nil
	}
	return container.Index
}

func LoadDates() []DateItem {
	url := apiBase + "/dates"
	var container DatesIndex
	if err := fetchAndCache(url, "data/dates.json", &container); err != nil {
		log.Println("LoadDates: returning empty dates due to error:", err)
		return nil
	}
	return container.Index
}

func LoadRelations() []RelationItem {
	url := apiBase + "/relation"
	var container RelationsIndex
	if err := fetchAndCache(url, "data/relations.json", &container); err != nil {
		log.Println("LoadRelations: returning empty relations due to error:", err)
		return nil
	}
	return container.Index
}
