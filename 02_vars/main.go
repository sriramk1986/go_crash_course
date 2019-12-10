package main

import "fmt"

func main() {
	var name string = "Brad"
	var age int = 37
	const isCool = true
	isCool = false

	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n", isCool)
}
