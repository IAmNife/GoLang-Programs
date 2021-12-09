package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("--------------------------")
	fmt.Println("Count Function: ", strings.Count("india", "i"))

	fmt.Println("--------------------------")
	fmt.Println("Equal Fold Function: ", strings.EqualFold("Cat", "cAt"))
	fmt.Println("Equal Fold Function: ", strings.EqualFold("India", "Indiana"))

}
