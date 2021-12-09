package main

import "fmt"

func linearSearch(dataList []int, key int) bool {
	for _, item := range dataList {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{12, 22, 43, 13, 35, 53, 64, 31, 54}
	fmt.Println(linearSearch(items, 12))
}
