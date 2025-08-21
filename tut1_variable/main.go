package main

import (
	"fmt"
	"learning/greetuser"
)

func main() {
	fmt.Println("____Gada_Electronics____")
	var name string = "Arpit Agrahari"
	var age int = 20
	const address = "Gola Bazar"
	phone := 9599056856
	var total_price float64 = 870.5
	greetuser.GreetUser(name)
	fmt.Println("____User Details____")
	greetuser.Userdetails(name, age, address, phone, total_price)

}
