package main

import (
	"fmt"
)

func determineGrade(score int) string {

	if score < 0 || score > 101 {
		return "Invalid score"
	} else if score >= 85 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 55 {
		return "C"
	} else if score >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	var score int
	fmt.Print("Enter the score: ")

	_, err := fmt.Scan(&score)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	grade := determineGrade(score)
	fmt.Println("The grade is: ", grade)
}
