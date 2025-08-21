package greetuser

import "fmt"

func GreetUser(name string) {
	fmt.Println("Hello " + name)
}
func Userdetails(name string, age int, address string, phone int, total_price float64) {
	fmt.Println("Name: " + name)
	fmt.Println("Age: ", age)
	fmt.Println("Address: ", address)
	fmt.Println("Phone: ", phone)
	fmt.Println("Total Price: ", total_price)
}
