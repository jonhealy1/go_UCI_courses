package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const partSize = 4

func sortInts(ints []int, c chan []int) {
	fmt.Println("Sorting: ", ints)
	sort.Ints(ints)
	c <- ints
}

func main() {

	var intsToSort, intsSorted []int
	c := make(chan []int, partSize)

	fmt.Println("enter sequence of integers to sort")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	sSlice := strings.Split(line, " ")

	if len(sSlice) < partSize {
		fmt.Println("Please enter minimum ", partSize, " ints")
		os.Exit(1)
	}

	for _, s := range sSlice {
		i, _ := strconv.Atoi(s)
		intsToSort = append(intsToSort, i)
	}

	go sortInts(intsToSort[0:len(intsToSort)/partSize], c)
	go sortInts(intsToSort[len(intsToSort)/partSize:len(intsToSort)/2], c)
	go sortInts(intsToSort[len(intsToSort)/2:len(intsToSort)*3/partSize], c)
	go sortInts(intsToSort[len(intsToSort)*3/partSize:len(intsToSort)], c)

	for i := 0; i < partSize; i++ {
		intsSorted = append(intsSorted, <-c...)
	}
	sort.Ints(intsSorted)
	fmt.Println("Sorted ints:", intsSorted)
}