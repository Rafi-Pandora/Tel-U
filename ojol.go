package main

import "strconv"

func Ojol(WaktuJam, WaktuMenit int, Jarak float64) string {
	const str string = "\nmaaf kami tidak dapat melayani pesanan anda"
	var tarifPerKm int
	waktu := strconv.Itoa(WaktuJam) + ":" + strconv.Itoa(WaktuMenit)

	if (WaktuJam >= 5 && WaktuJam <= 9) || (WaktuJam >= 16 && WaktuJam <= 19) {
		if Jarak > 0 && Jarak <= 10 {
			tarifPerKm = 5000
		} else if Jarak > 10 && Jarak <= 20 {
			tarifPerKm = 4500
		} else {
			return str
		}
	} else if WaktuJam <= 10 && WaktuJam >= 5 {
		if Jarak > 0 && Jarak <= 20 {
			tarifPerKm = 4000
		} else {
			return str
		}
	} else {
		return str
	}

	return "\npemesanan pada waktu " + waktu + ", jarak " + strconv.Itoa(int(Jarak)) + " adalah sebesar: " + strconv.Itoa(tarifPerKm*int(Jarak))
}
