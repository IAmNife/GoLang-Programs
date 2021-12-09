package main

import "fmt"

func main() {

	var radius float64
	const PI float64 = 3.14
	var area float64
	var circumference float64

	fmt.Println("Enter radius of Circle : ")
	fmt.Scan(&radius)
	
	area = PI * radius * radius
	fmt.Println("Area of the circle is ", area)

	circumference = 2 * PI * radius
	fmt.Println("Circumference of the circle is",circumference)

}