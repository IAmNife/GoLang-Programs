package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("--------------------------")
	fmt.Println("Contains function: ", strings.Contains("India", "ia"))
	fmt.Println("Contains function: ", strings.Contains("India", "ma"))

	fmt.Println("--------------------------")
	fmt.Println("Has Prefix: ", strings.HasPrefix("India", "In"))
	fmt.Println("Has Prefix: ", strings.HasPrefix("India", "Na"))

	fmt.Println("--------------------------")
	fmt.Println("Has Suffix: ", strings.HasSuffix("India", "ia"))
	fmt.Println("Has Suffix: ", strings.HasSuffix("India", "la"))

	fmt.Println("--------------------------")
	fmt.Println("Join Function: ", strings.Join([]string{"Karnataka", "Andra Pradesh", "Tamil Nadu"}, "-"))
	fmt.Println("Join Function: ", strings.Join([]string{"Karnataka", "Andra Pradesh", "Tamil Nadu"}, " "))

	fmt.Println("--------------------------")
	fmt.Println("Repeat Function: ", strings.Repeat("India", 3))

	fmt.Println("--------------------------")
	fmt.Println("Replace Function: ", strings.Replace("India", "I", "N", 1))
	fmt.Println("Replace Function: ", strings.Replace("India", "I", "N", 2))
	fmt.Println("Replace Function: ", strings.Replace("India", "I", "N", 3))
	fmt.Println("Replace Function: ", strings.Replace("India", "dia", "idn", 1))

	fmt.Println("--------------------------")
	fmt.Println("Split Function: ", strings.Split("Karnataka-Andra Pradesh-Tamil Nadu", "-"))

	fmt.Println("--------------------------")
	fmt.Println("To Lower Function", strings.ToLower("India"))

	fmt.Println("--------------------------")
	fmt.Println("To Upper Function", strings.ToUpper("India"))
}
