package models

import "strings"

func FindArtist(id int, artists []Artist) *Artist {
	for _, a := range artists {
		if a.ID == id {
			return &a
		}
	}
	return nil
}

func FindLocations(id int, locs []LocationItem) []string {
	for _, l := range locs {
		if l.ID == id {
			return l.Locations
		}
	}
	return []string{}
}

func CleanDates(dates []string) []string {
	cleaned := make([]string, len(dates))
	for i, date := range dates {
		cleaned[i] = strings.TrimPrefix(strings.TrimSpace(date), "*")
	}
	return cleaned
}

func FindDates(id int, dates []DateItem) []string {
	for _, d := range dates {
		if d.ID == id {
			return CleanDates(d.Dates)
		}
	}
	return []string{}
}

func FindRelation(id int, rels []RelationItem) map[string][]string {
	for _, r := range rels {
		if r.ID == id {
			cleaned := make(map[string][]string)
			for city, dates := range r.DatesLocations {
				cleaned[city] = CleanDates(dates)
			}
			return cleaned
		}
	}
	return map[string][]string{}
}
