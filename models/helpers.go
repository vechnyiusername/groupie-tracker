package models

import "strings"

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

// Удалить звёздочки из дат
func CleanDates(dates []string) []string {
	cleaned := make([]string, len(dates))
	for i, date := range dates {
		// Удаляем звёздочку в начале строки, если она есть (с учётом пробелов)
		cleaned[i] = strings.TrimPrefix(strings.TrimSpace(date), "*")
	}
	return cleaned
}

// Найти даты концертов по ID артиста
func FindDates(id int, dates []DateItem) []string {
	for _, d := range dates {
		if d.ID == id {
			return CleanDates(d.Dates)
		}
	}
	return []string{}
}

// Найти связи (город → даты) по ID артиста
func FindRelation(id int, rels []RelationItem) map[string][]string {
	for _, r := range rels {
		if r.ID == id {
			// Очищаем даты от звёздочек в каждой записи
			cleaned := make(map[string][]string)
			for city, dates := range r.DatesLocations {
				cleaned[city] = CleanDates(dates)
			}
			return cleaned
		}
	}
	return map[string][]string{}
}
