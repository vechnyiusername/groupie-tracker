package models

type ArtistFull struct {
	Artist
	Locations      []string
	Dates          []string
	DatesLocations map[string][]string
}
