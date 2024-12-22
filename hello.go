package main

import "fmt"

func changeForV2() int {
	fmt.Print("Change for v2")
	return 2
}

func main() {
	x := 5

	fmt.Println("Hello, World!", x)
	fmt.Println("Hello, World!", changeForV2())
}
