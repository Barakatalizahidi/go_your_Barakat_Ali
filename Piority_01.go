package main

import (
	"fmt"
)

const Phi = 3.14

func calculateVolume(radius float64, height float64) float64 {
	return Phi * radius * radius * height
}

func main() {
	var radius float64
	var height float64

	fmt.Println("Enter the radius of the tube: ")
	fmt.Scan(&radius)

	fmt.Printf("Enter the height of the tube: ")
	fmt.Scan(&height)

	volume := calculateVolume(radius, height)

	fmt.Printf("The volume of the tube is: %.2f\n", volume)

}
