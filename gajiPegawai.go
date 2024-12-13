package main

import "strconv"

func GajiPegawai(Jabatan string, MasaKerja uint, JumlahAnak uint) string {
	var gajiPokok uint
	var tunjangan uint = 0

	if JumlahAnak > 3 {
		JumlahAnak = 3
	}

	switch Jabatan {
	case "Staf":
		if MasaKerja < 5 {
			gajiPokok = 4000
		} else if MasaKerja > 5 && MasaKerja < 10 {
			gajiPokok = 4000
			tunjangan = 100
		} else {
			gajiPokok = 5000
			tunjangan = 100
		}
	case "Direktur":
		gajiPokok = 20000
		tunjangan = 500
	case "Manager":
		tunjangan = 300
		if MasaKerja > 10 {
			gajiPokok = 10000
		} else {
			gajiPokok = 8500
		}
	default:
		return "input salah!, input yang diizinkan: Staf|Direktur|Manager (string), Masa Kerja (uint), Jumlah anak (uint)"
	}

	return strconv.Itoa(int(gajiPokok)) + " + " + strconv.Itoa(int(JumlahAnak*tunjangan)) + " = " + strconv.Itoa(int(gajiPokok+(JumlahAnak*tunjangan)))
}
