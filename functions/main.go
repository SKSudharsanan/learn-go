package main

import (
	"errors"
	"fmt"
)

// basic function
func add(x int, y int) int {
	return x + y
}

// return multiple values
func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("cannot be divided by zero")
	}
	return x / y, nil
}

// return named values
func stats(number []float64) (mean float64, sum float64) {
	for _, num := range number {
		sum += num
	}
	mean = sum / float64(len(number))
	return
}

func calulateSum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func compute(fn func(float64, float64) float64, a float64, b float64) float64 {
	return fn(a, b)
}

func multiply(x, y float64) float64 {
	return x * y
}

func makeMultiplier(factor float64) func(float64) float64 {
	return func(input float64) float64 {
		return input * factor
	}
}

func main() {

	//basic function
	a := 3
	b := 10
	fmt.Println(add(a, b))

	//multiple values function
	x := 5.0
	y := 25.0
	fmt.Println(divide(x, y))

	//named values return
	numbers := []float64{1.0, 76.21, 24.8, 43.6, 76.9}
	mean, sum := stats(numbers)
	fmt.Println("Mean:", mean, "Sum:", sum)

	//variadic functions
	//If you want to pass a variable number of arguments of the same type to a function, you can use a variadic function. The function sum below can take any number of int parameters
	fmt.Println(calulateSum(4, 3, 5, 7, 8))

	//functions as type
	var myAdd func(int, int) int
	myAdd = add
	result := myAdd(10, 5)
	fmt.Println("Result:", result)

	//functions as arguments
	newresult := compute(multiply, 3, 4)
	fmt.Println("Multiply 3 and 4:", newresult)

	//functions as return types useful for factory functions
	doubler := makeMultiplier(2)
	tripler := makeMultiplier(3)

	fmt.Println(doubler(5))
	fmt.Println(tripler(5))
}
