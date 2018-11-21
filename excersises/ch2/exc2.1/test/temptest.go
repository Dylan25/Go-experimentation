package main

import "Go-experimentation/excersises/ch2/exc2.1"

func main() {
	tempC := tempconv.Celsius(20)
	tempF := tempconv.CToF(tempC)
	tempK := tempconv.FToK(tempF)

	println(tempC.String())
	println(tempF.String())
	println(tempK.String())
}
