package models

// FindArtist returns the artist by ID
func FindArtist(id int, artists []Artist) *Artist {
	for _, a := range artists {
		if a.ID == id {
			return &a
		}
	}
	return nil
}

// FindLocations returns the locations by artist ID
func FindLocations(id int, locs []Locations) []string {
	for _, l := range locs {
		if l.ID == id {
			return l.Locations
		}
	}
	return nil
}

// FindDates returns the concert dates by artist ID
func FindDates(id int, dates []Dates) []string {
	for _, d := range dates {
		if d.ID == id {
			return d.Dates
		}
	}
	return nil
}

// FindRelation returns map city -> dates
func FindRelation(id int, rels []Relation) map[string][]string {
	for _, r := range rels {
		if r.Index == id {
			return r.DatesLocations
		}
	}
	return nil
}
