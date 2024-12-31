package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// fungsi ini digunakan untuk mendeklarasikan database
func InsialisasiDatabase() (*sql.DB, error) {
	const username string = "admin123"
	const password string = "admin123"

	//untuk menampung nilai dari username dan password ditambah dengan URI database
	dns := username + ":" + password + "@tcp(localhost:3306)/mahasiswa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dns)

	// nill = null
	if err != nil {
		return nil, err
	}

	//cek koneksi
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func LihatSemuaUserDiDatabase(db *sql.DB) ([]string, []string) {
	var namaArr = []string{}
	var nimArr = []string{}

	//coba query ke database
	rows, err := db.Query("SELECT nama, nim FROM mahasiswa")
	if err != nil {
		log.Fatal("gagal menjalankan queri ", err)
	}
	defer rows.Close()

	// Perulangan
	for rows.Next() {
		var name string
		var nim string

		// buat narik data dari database
		if err := rows.Scan(&name, &nim); err != nil {
			log.Fatal(err)
		}
		namaArr = append(namaArr, name)
		nimArr = append(nimArr, nim)
	}

	//cetak jika ada error di loop
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return namaArr, nimArr
}

func TambahUserDatabase(db *sql.DB, nama string, nim uint) error {
	query := `INSERT INTO mahasiswa (nama, nim) VALUES (?, ?)`
	_, err := db.Exec(query, nama, nim)
	return err
}

func HapusUserDatabase(db *sql.DB, nama string, nim string) error {
	var query string
	var params []interface{}

	if nama != "" {
		query = `DELETE FROM mahasiswa WHERE nama = ?`
		params = append(params, nama)
	} else if nim != "" {
		query = `DELETE FROM mahasiswa WHERE nim = ?`
		params = append(params, nim)
	} else {
		log.Println("Nama atau nih harus diisi")
	}

	_, err := db.Exec(query, params...)
	return err
}

func AksesDatabaseAdmin(db *sql.DB, nama string, pass string) (string, error) {
	var user string
	query := "SELECT username FROM admin WHERE username = ? AND  password = ?"
	err := db.QueryRow(query, nama, pass).Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			return "", err
		}
	}
	return user, nil
}
