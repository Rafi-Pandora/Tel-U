// package atau class yang berisi fungsi pendukung
package main

import "fmt"

func InputInt(pesan string) int {
	var i int
	println(pesan)
	_, err := fmt.Scanln(&i)
	if err != nil {
		println(err)
		return 0
	}
	return i
}

func InputFloat(pesan string) float64 {
	var i float64
	println(pesan)
	_, err := fmt.Scanln(&i)
	if err != nil {
		println(err)
		return 0
	}
	return i
}

func InputStr(pesan string) string {
	var i string
	println(pesan)
	_, err := fmt.Scanln(&i)
	if err != nil {
		println(err)
		return ""
	}
	return i
}

func InputBool(pesan string) bool {
	var i bool
	println(pesan)
	_, err := fmt.Scanln(&i)
	if err != nil {
		println(err)
		return false
	}
	return i
}
