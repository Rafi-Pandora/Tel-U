package main

func soal2(input uint) string {
	const kelipatan3 string = "kelipatan 3"
	const kelipatan5 string = "kelipatan 5"

	if input%3 == 0 && input%5 == 0 {
		return kelipatan3 + "\n" + kelipatan5
	} else if input%5 == 0 {
		return kelipatan5
	} else if input%3 == 0 {
		return kelipatan3
	} else {
		return ""
	}
}
