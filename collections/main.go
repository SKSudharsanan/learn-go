package main

import (
	"fmt"
)

func main() {
	//arrays
	var arr [5]int = [5]int{12, 30, 90, 87, 21}
	for index, value := range arr {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//slices - they are arrays of dynamic size
	var slice []int = make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 25)
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	//maps - key-value pairs
	var dict map[string]int = make(map[string]int)
	dict["one"] = 1
	dict["two"] = 2
	fmt.Println(dict["one"])
}
