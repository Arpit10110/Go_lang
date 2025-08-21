package main

import "fmt"

func GreetUser(name string) {
	fmt.Println("Hello " + name)
}

func sum(a int, b int) int {
	return a + b
}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {
	fmt.Println("---Learning Functions-----")
	GreetUser("Arpit")
	Add := sum(5, 7)
	fmt.Println("Sum of 5+7 is", Add)
	Multi := multiply(5, 7)
	fmt.Println("Multiply of 5x7 is", Multi)
}
