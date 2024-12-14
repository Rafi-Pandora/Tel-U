package main

func main() {
	println(Ktp(uint(InputInt("masukan usia :")), InputBool("memiliki Kartu keluarga? True|False")))
	println(TigaBialangan(InputInt("\n\nmasukan angka a :"), InputInt("masukan angka b :"), InputInt("masukan angka c :")))
	println(ManagerEPL())
	println(GajiPegawai(InputStr("\n\nJabatan :"), uint(InputInt("Masa Kerja :")), uint(InputInt("Jumlah Anak :"))))
	println(Ojol(InputInt("Masukan jam:"), InputInt("Masukan menit:"), float64(InputInt("Masukan Jarak:"))))
}
