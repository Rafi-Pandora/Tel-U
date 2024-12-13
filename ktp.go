package main

func Ktp(usia uint, KartuKeluaraga bool) string {
	if usia >= 17 && KartuKeluaraga {
		return "bisa membuat ktp"
	} else {
		return "belum bisa membuat ktp"
	}
}
