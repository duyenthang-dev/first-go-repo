package main

import (
	scopeone "example/hello/scope/scope-one"
	"fmt"
)

var one = "One"

func main() {

	one := one
	fmt.Println(one)
	myFunc()
}

func myFunc() {
	var one = "The number one"
	fmt.Println(one)
	fmt.Println(scopeone.PublicOne)
}
