package main

import (
	"errors"
	"fmt"
)

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()
	panic("something bad happened")
}

func calculateGrade(score int) (string, error) {
	if score < 0 || score > 100 {
		return "", errors.New("invalid score")
	}
	switch {
	case score >= 90:
		return "A", nil
	case score >= 80:
		return "B", nil
	case score >= 70:
		return "C", nil
	case score >= 60:
		return "D", nil
	default:
		return "F", nil
	}
}

func main() {
	grade, err := calculateGrade(105) // Trying an invalid score
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Grade:", grade)
	}
	riskyFunction()
	fmt.Println("Program continues...")

}
