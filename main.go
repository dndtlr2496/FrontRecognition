package main

import "github.com/dndtlr2496/FrontRecognition/module"

func main() {
	aX := 1
	aY := 1
	frontVector := [2]int{0, 1}
	bX := 4
	bY := 4
	module.Start(aX, aY, frontVector, bX, bY)
}
