package utils

import (
	"html/template"
	"server/models"
	"strconv"
)

// Untuk halaman list atau main
type ViewData struct {
	Name      string
	FormLabel string
	Films     []*models.Film
}

// Untuk form tambah/edit
type FilmFormData struct {
	FormTitle   string
	FormAction  string
	SubmitLabel string
	Film        *models.Film
}

var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
}

// helper untuk parsing template dengan funcMap
func ParseTemplate(filename string, filepath string) (*template.Template, error) {
	return template.New(filename).Funcs(funcMap).ParseFiles(filepath)
}

// helper cari index film dari pointer tanpa break (gunakan indeks terakhir yg cocok)
func FindFilmIndex(data []*models.Film, target *models.Film) int {
	idx := -1
	for i, f := range data {
		if f == target {
			idx = i
		}
	}
	return idx
}

// helper filter films non-nil
func FilterNonNilFilms(data []*models.Film) []*models.Film {
	films := make([]*models.Film, 0, len(data))
	for _, f := range data {
		if f != nil {
			films = append(films, f)
		}
	}
	return films
}

func RatingFinal(rating1, rating2 string) (float64, error) {
	parseRating := func(s string) (float64, error) {
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0, err
		}
		return f, nil
	}

	r1, err1 := parseRating(rating1)
	r2, err2 := parseRating(rating2)

	if err1 != nil {
		return 0, err1
	} else if err2 != nil {
		return 0, err2
	}

	switch {
	case r1 == 0 && r2 == 0:
		return 0, nil
	case r1 == 0:
		return r2, nil
	case r2 == 0:
		return r1, nil
	default:
		if r1 > r2 {
			return r1, nil
		}
		return r2, nil
	}
}
