package main

import "fmt"
import s "strings"

func main() {
	var input string
	var inputLow string
	fmt.Printf("Enter a string please:\n")
	n, err := fmt.Scan(&input)

	if n != 1 || err != nil {
		fmt.Println("Error.")
		return
	}
    inputLow = s.ToLower(input)
	lastChar := inputLow[len(input)-1:]
	firstChar := inputLow[:1]

	if lastChar == "n" && firstChar == "i" && s.ContainsRune(inputLow, 97) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
	
	//fmt.Println("Answer: ", lastChar, firstChar)
}
