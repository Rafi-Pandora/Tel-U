package main

func soal3(input0 uint, input1 uint, input2 uint) string {

	switch {
	case input0 == input1 && input0 == input2:
		return "segitiga sama sisi"

	case input0 == input1 || input0 == input2 || input1 == input2:
		return "segitiga sama kaki"

	default:
		return "segitiga sembarang"
	}
}
