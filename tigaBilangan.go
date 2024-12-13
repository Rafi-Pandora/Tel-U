package main

import "strconv"

func TigaBialangan(a, b, c int) string {
	switch {
	case a <= c && c <= b:
		b, c = c, b
		/*
			menukar nilai
			nilai a tetap
			b = c
			c = b
		*/
	case b <= a && a <= c:
		a, b = b, a
		/*
			menukar nilai
			a = b
			b = a
			nilai c tetap
		*/
	case b <= c && c <= a:
		a, b, c = b, c, a
		/*
			menukar nilai
			a = b
			b = c
			c = a
		*/
	case c <= a && a <= b:
		a, b, c = c, a, b
		/*
			menukar nilai
			a = c
			b = a
			c = b
		*/
	case c <= b && b <= a:
		a, c = c, a
		/*
			menukar nilai
			a = c
			nilai b tetap
			c = b
		*/
	case a <= b && b <= c:
		//tidak melakukan apa-apa karena nilai kembalian sama
	}
	return strconv.Itoa(a) + " " + strconv.Itoa(b) + " " + strconv.Itoa(c)
}
