package models

// Struct untuk data yang akan dikirim ke template
type Film struct {
	Judul  string
	Genre  string
	Rating float32
	Status bool
}

func SeedFilms() []*Film {
	return []*Film{
		{Judul: "Avengers", Genre: "Action", Rating: 8.5, Status: true},
		{Judul: "Interstellar", Genre: "Science Fiction / Sci-Fi", Rating: 9.0, Status: false},
		{Judul: "The Dark Knight", Genre: "Action", Rating: 9.0, Status: true},
		{Judul: "Inception", Genre: "Science Fiction / Sci-Fi", Rating: 8.8, Status: true},
		{Judul: "Parasite", Genre: "Drama", Rating: 8.6, Status: false},
		{Judul: "The Godfather", Genre: "Crime", Rating: 9.2, Status: true},
		{Judul: "Pulp Fiction", Genre: "Crime", Rating: 8.9, Status: false},
		{Judul: "Forrest Gump", Genre: "Drama", Rating: 8.8, Status: true},
		{Judul: "Spirited Away", Genre: "Animation", Rating: 8.6, Status: false},
		{Judul: "Whiplash", Genre: "Drama", Rating: 8.5, Status: true},
		{Judul: "Coco", Genre: "Animation", Rating: 8.4, Status: true},
		{Judul: "Your Name", Genre: "Romance", Rating: 8.4, Status: false},
		{Judul: "The Matrix", Genre: "Science Fiction / Sci-Fi", Rating: 8.7, Status: true},
		{Judul: "Joker", Genre: "Drama", Rating: 8.4, Status: false},
		{Judul: "Gladiator", Genre: "Action", Rating: 8.5, Status: true},
		{Judul: "Up", Genre: "Animation", Rating: 8.2, Status: false},
		{Judul: "The Lion King", Genre: "Animation", Rating: 8.5, Status: true},
		{Judul: "Titanic", Genre: "Romance", Rating: 7.8, Status: true},
		{Judul: "Frozen", Genre: "Animation", Rating: 7.5, Status: false},
		{Judul: "Shutter Island", Genre: "Thriller", Rating: 8.1, Status: true},
		{Judul: "A Beautiful Mind", Genre: "Drama", Rating: 8.2, Status: false},
		{Judul: "Fight Club", Genre: "Drama", Rating: 8.8, Status: true},
		{Judul: "La La Land", Genre: "Musical", Rating: 8.0, Status: false},
		{Judul: "Dune", Genre: "Science Fiction / Sci-Fi", Rating: 8.1, Status: true},
		{Judul: "Blade Runner 2049", Genre: "Science Fiction / Sci-Fi", Rating: 8.0, Status: false},
		{Judul: "The Shawshank Redemption", Genre: "Drama", Rating: 9.3, Status: true},
		{Judul: "The Prestige", Genre: "Mystery", Rating: 8.5, Status: true},
		{Judul: "The Social Network", Genre: "Biography", Rating: 7.7, Status: false},
		{Judul: "Interstella 5555", Genre: "Music", Rating: 7.8, Status: false},
		{Judul: "The Pianist", Genre: "Biography", Rating: 8.5, Status: true},
	}
}
