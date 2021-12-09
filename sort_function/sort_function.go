package main

import (
	"fmt"
	"sort"
)

func main() {

	num := []int{50, 90, 30, 10, 50}
	fmt.Println(num)
	if sort.IntsAreSorted(num) == false {
		sort.Ints(num)
	}
	fmt.Println(num)

	text := []string{"Nife", "Niroop", "Arsenal", "RCB"}
	fmt.Println(text)
	if sort.StringsAreSorted(text) == false {
		sort.Strings(text)
	}
	fmt.Println(text)

}
