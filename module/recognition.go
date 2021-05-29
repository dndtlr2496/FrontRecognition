package module

import "fmt"

func CalcVector(aX int, aY int, bX int, bY int) [2]int {
	return [2]int{bX - aX, bY - aY}
}

func Dot(frontVector [2]int, b [2]int) bool {
	return frontVector[0]*b[0]+frontVector[1]*b[1] > 0
}

func Start(aX int, aY int, frontVector [2]int, bX int, bY int) int {
	b := CalcVector(aX, aY, bX, bY)

	if Dot(frontVector, b) {
		fmt.Println("Front")
		return 1
	}
	fmt.Println("Not Front")
	return 1
}
