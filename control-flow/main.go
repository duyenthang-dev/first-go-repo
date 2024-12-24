package main

import "math/rand"

// go has init function that runs before main
func init() {
	println("Control Flow")
}

func main() {
	var x int = 5

	if x > 5 {
		println("x is greater than 5")
	} else if x < 5 {
		println("x is less than 5")
	} else {
		println("x is equal to 5")
	}

	// statement can precede condition in if statement
	y := 40
	if z := 2 * rand.Intn(y); y > z {
		println("y is greater than 2 times z")
	} else {
		println("y is less than or equal to 2 times z")
	}

	// comma ok idiom

}
