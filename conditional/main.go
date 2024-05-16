package main

import "fmt"

// type switch
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	//if-else - check odd or even
	var n int = 15
	if n%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}

	//compound operations
	age := 25
	income := 50000
	if age > 18 && income > 40000 {
		fmt.Println("Eligible for credit")
	}

	//switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	//switch with no expression
	x := 42
	switch {
	case x < 20:
		fmt.Println("x is less than 20")
	case x > 20:
		fmt.Println("x is greater than 20")
	case x == 20:
		fmt.Println("x is equal to 20")
	default:
		fmt.Println("x is not a number")
	}

	do(21)
	do("hello")
	do(true)

}
