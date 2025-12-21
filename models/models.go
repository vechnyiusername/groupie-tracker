package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type LocationsIndex struct {
	Index []LocationItem `json:"index"`
}

type LocationItem struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DatesIndex struct {
	Index []DateItem `json:"index"`
}

type DateItem struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationsIndex struct {
	Index []RelationItem `json:"index"`
}

type RelationItem struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type ArtistFull struct {
	Artist
	Locations      []string
	Dates          []string
	DatesLocations map[string][]string
}
