package main

import "fmt"

func main() {

	var first, second string

	fmt.Print("Enter the first string ")
	fmt.Scan(&first)
	fmt.Print("Enter the second string ")
	fmt.Scan(&second)

	fmt.Print("Combined string is ", first+second)
}
