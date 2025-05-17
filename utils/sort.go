package utils

import (
	"server/models"
)

func InsertionSortByRating(films []*models.Film, ascending bool) {
	for i := 1; i < len(films); i++ {
		key := films[i]
		j := i - 1

		// Sorting in ascending order
		if ascending {
			for j >= 0 && films[j].Rating > key.Rating {
				films[j+1] = films[j]
				j--
			}
		} else { // Sorting in descending order
			for j >= 0 && films[j].Rating < key.Rating {
				films[j+1] = films[j]
				j--
			}
		}
		films[j+1] = key
	}
}

func InsertionSortByJudul(films []*models.Film, ascending bool) {
	for i := 1; i < len(films); i++ {
		key := films[i]
		j := i - 1

		// Sorting in ascending order
		if ascending {
			for j >= 0 && films[j].Judul > key.Judul {
				films[j+1] = films[j]
				j--
			}
		} else { // Sorting in descending order
			for j >= 0 && films[j].Judul < key.Judul {
				films[j+1] = films[j]
				j--
			}
		}
		films[j+1] = key
	}
}
