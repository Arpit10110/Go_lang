package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Println("Enter your name: ")
	// it can not read after space
	// fmt.Scan(&name)
	// fmt.Println("Hello", name)
	// to read after space we can use bufio
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("Hello", name)
}
