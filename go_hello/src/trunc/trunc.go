package main

import "fmt"

func main() {
	var numeral float64
	fmt.Printf("Enter a floating point number:\n")
	fmt.Scan(&numeral)

	fmt.Println("Answer: ", int(numeral))
}
