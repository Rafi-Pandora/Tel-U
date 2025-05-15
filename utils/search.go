package utils

import (
	"server/models"
)

func BinarySearchByGenre(films []*models.Film, target string) (int, *models.Film) {
	low, high := 0, len(films)-1
	for low <= high {
		mid := (low + high) / 2
		if films[mid].Genre == target {
			return mid, films[mid]
		} else if films[mid].Genre < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, nil
}

func BinarySearchByJudul(films []*models.Film, target string) (int, *models.Film) {
	low, high := 0, len(films)-1
	for low <= high {
		mid := (low + high) / 2
		if films[mid].Judul == target {
			return mid, films[mid]
		} else if films[mid].Judul < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, nil
}
