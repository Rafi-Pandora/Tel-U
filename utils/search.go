package utils

import (
	"server/models"
	"strconv"
	"strings"
)

func SequentialSearchByGenre(films []*models.Film, target string) []*models.Film {
	var result []*models.Film
	for _, film := range films {
		if film.Genre == target {
			result = append(result, film)
		}
	}
	return result
}

func SequentialSearchByRating(films []*models.Film, target string) ([]*models.Film, error) {
	var result []*models.Film
	targetFloat, err := strconv.ParseFloat(target, 32)

	min := float32(targetFloat)
	max := min + 1

	if err != nil {
		return nil, err
	}

	for _, film := range films {
		if film.Rating >= min && film.Rating < max {
			result = append(result, film)
		}
	}
	return result, nil
}

func BinarySearchByJudul(films []*models.Film, target string) (int, *models.Film) {
	var lowerCaseJudul, lowerCaseTarget string

	lowerCaseTarget = strings.ToLower(target)
	low, high := 0, len(films)-1

	for low <= high {
		mid := (low + high) / 2
		lowerCaseJudul = strings.ToLower(films[mid].Judul)

		if lowerCaseJudul == lowerCaseTarget {
			return mid, films[mid]
		} else if lowerCaseJudul < lowerCaseTarget {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1, nil
}
