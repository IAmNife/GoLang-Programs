package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{23, 43, 4, 345, 56, 12, -34}
	fmt.Println("Slice of Integers before sorting ", intSlice)
	sort.Ints(intSlice)
	fmt.Println("Slice of Integers after sortng", intSlice)
}
