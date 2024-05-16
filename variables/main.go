package main

import "fmt"

func main() {
	//var keyword is used to declare variables
	var age = 30
	fmt.Println(age)
	// := is the shorthand to declare a variable
	name := "sudharsanan"
	fmt.Println(name)
	//datatypes
	//integers - int, uint, int8, int16, int32, int64, uint8, uint16,uint32, uint64
	var a int = 12
	fmt.Println(a)
	var b uint = 30
	fmt.Println(b)
	var aa int8 = -29
	fmt.Println(aa)
	var ab uint8 = 32
	fmt.Println(ab)

	//float32, float64
	var afloat float32 = 23.4
	fmt.Println(afloat)

	//string
	var msg string = "learn easy code"
	fmt.Println(msg)

	//bool
	var isalive bool = true
	fmt.Println(isalive)

}
