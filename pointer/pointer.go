package main

import "fmt"

func doSomething(a *int) {
	fmt.Println("a (value) in function :", *a)
	*a = 20
}

func main() {
	var a int = 10
	var pA *int = &a

	fmt.Println("a (value) before function :", a)
	doSomething(&a)
	fmt.Println("a (value) after function  :", a)
	fmt.Println("pA (value) :", pA)

	sliceDemo()
	mapDemo()
}

func sliceDemo() {
	var animals []string = []string{"cat", "dog", "bird", "fish", "elephant"}
	fmt.Println("animals: ", animals)

	DeleteFromSlice(animals, 2)

	fmt.Println("animals: ", animals)
}

func DeleteFromSlice(slice []string, index int) []string {
	var left = slice[:index]
	var right = slice[index+1:]
	fmt.Println("left: ", left)
	fmt.Println("right: ", right)
	return append(slice[:index], slice[index+1:]...)
}

func mapDemo() {
	fmt.Println("--- Map Demo ---")
	var myMap map[int]string = make(map[int]string)
	myMap[1] = "One"
	myMap[2] = "Two"
	myMap[3] = "Three"
	fmt.Println("myMap: ", myMap)
	delete(myMap, 2)

	fmt.Println("myMap: ", myMap)
}
