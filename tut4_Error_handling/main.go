package main

import "fmt"

// func divide(a, b float64) (float64, string) {
// 	if b == 0 {
// 		return 0, "cannot divide by zero"
// 	} else {
// 		return a / b, "nil"
// 	}
// }

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	} else {
		return a / b, nil
	}
}

func main() {
	fmt.Println("---Learning Error Handling-----")
	result, _ := divide(10, 0)
	fmt.Println("Result is", result)
	// result, err := divide(10, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Result is", result)
	// }
}
