package main

import (
	"fmt"
	"time"
	"strconv"
	"sync"
)

var isit int

func main() {
	start := time.Now()

	var a,b,c,d int
	var wg sync.WaitGroup
	wg.Add(4)
	isit = 0
	s := userInput()
	/*s := []int{7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,345655,
		7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,
		//345655,7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,
		//345655,7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,345655,
		//7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,345655,7,6,5,4,3,2,1,
		//343,23,5,345435,3453,434,34,34,34,345655,45,5,5,43,467,7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,345655,
		//7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,
		//345655,7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,
		//345655,7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,345655,
		//7,6,5,4,3,2,1,343,23,5,345435,3453,434,34,34,34,345655,7,6,5,4,3,2,1,
		343,23,5,345435,3453,434,34,34,34,345655,45,5,5,43,467,8}*/

	a = len(s)/4
	b = len(s)/4 + a
	c = len(s)/4 + b + c
	d = len(s)

	s1 := arrange(0,a,s)
	s2 := arrange(a,b,s)
	s3 := arrange(b,c,s)
	s4 := arrange(c,d,s)

	go sortGO(s1, &wg)
	go sortGO(s2, &wg)
	go sortGO(s3, &wg)
	go sortGO(s4, &wg)

/*	go sort(s1)
	go sort(s2)
	go sort(s3)
	go sort(s4) */
/*
	sort(s1)
	sort(s2)
	sort(s3)
	sort(s4)
*/	
	wg.Wait()

	f1 := append(s1, s2...)
	f2 := append(s3, s4...)
	f3 := append(f1, f2...)
	sort(f3)

    elapsed := time.Since(start)
    fmt.Println("Sorting took %s", elapsed)
}

func arrange(a int, b int, s[] int) []int {
	x := s[a:b]
	return x
}

func userInput() []int {
	var numeral string
	loop := true
	s := make([]int, 0, 3)
	for loop {
		fmt.Println("Please input an integer or 'X' to quit.")
		fmt.Scan(&numeral)
		if(numeral == "X"){
		  break
		} else {
		  i1, err := strconv.Atoi(numeral)
		  if err != nil {
			fmt.Println(i1)
		  }
		  s = append(s, i1)
		}
	}
	fmt.Println(s)
	return s
}

func sortGO(s[] int, wg *sync.WaitGroup) []int {
	defer wg.Done()
	swapped := true
	isit = isit + 1
	fmt.Println("Go routine unsorted:", isit, "is:", s)
	for swapped {
		swapped = false
		for i:=0; i<len(s)-1; i++{
			if s[i] > s[i+1] {
				temp := s[i+1]
				s[i+1] = s[i]
				s[i] = temp
				swapped = true
			}
		}
	}
	
	fmt.Println("Go routine sorted:", s)
	return s
}

func sort(s[] int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i:=0; i<len(s)-1; i++{
			if s[i] > s[i+1] {
				temp := s[i+1]
				s[i+1] = s[i]
				s[i] = temp
				swapped = true
			}
		}
	}
	isit = isit + 1
	fmt.Println("Final routine:", isit, "is:", s)
	return s
}
