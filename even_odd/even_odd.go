package main

import "fmt"

func main() {

	fmt.Println("Enter an integer")
	var n int
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "is an even number")
	} else {
		fmt.Println(n, "is an odd number")
	}

}
