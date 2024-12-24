package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	var z int = add(10, 20)
	fmt.Println("z: ", z)

	var y int = sum(1, 2, 3, 4, 5)
	fmt.Println("y: ", y)

	var p Person = Person{
		Name: "John Doe",
		Age:  30,
		ID:   "1234",
	}
	p.SayHello()
	fmt.Println(p)
}

type Person struct {
	Name string
	Age  int
	ID   string
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old)", p.Name, p.Age)
}

func (p Person) SayHello() {
	fmt.Printf("Hello, my name is %v\n", p.Name)
}
