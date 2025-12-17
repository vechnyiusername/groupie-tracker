package models

// Найти артиста по ID
func FindArtist(id int, artists []Artist) *Artist {
	for _, a := range artists {
		if a.ID == id {
			return &a
		}
	}
	return nil
}

// Найти локации по ID артиста
func FindLocations(id int, locs []LocationItem) []string {
	for _, l := range locs {
		if l.ID == id {
			return l.Locations
		}
	}
	return []string{}
}

// Найти даты концертов по ID артиста
func FindDates(id int, dates []DateItem) []string {
	for _, d := range dates {
		if d.ID == id {
			return d.Dates
		}
	}
	return []string{}
}

// Найти связи (город → даты) по ID артиста
func FindRelation(id int, rels []RelationItem) map[string][]string {
	for _, r := range rels {
		if r.ID == id {
			return r.DatesLocations
		}
	}
	return map[string][]string{}
}
