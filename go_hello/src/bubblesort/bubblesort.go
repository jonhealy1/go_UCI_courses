package main

import (
  "fmt"
  "strconv"
)

func main() {
  var numerals string
  slice1 := make([]int, 0, 3)
  fmt.Printf("Please enter an integer (10 max). To exit enter 'X' :\n")
  
  for j := 0; j < 10; j++ {
    fmt.Scan(&numerals)
    if(numerals == "X"){
      break
    } else {
      i1, err := strconv.Atoi(numerals)
      if err != nil {
        fmt.Println(i1)
      }
	    
      slice1 = append(slice1, i1)
    }
  }
  bubbleSort(slice1)
  fmt.Println(slice1)
}

func bubbleSort(s []int) {
  //swap(s, 1)
  for i := 0; i < len(s)-1; i++ {
    for j := 0; j < len(s)-i-1; j++ {
      swap(s,j)
    }
  }
}

func swap(s []int, i int){
  if s[i] > s[i+1] {
    var temp int
    temp = s[i]
    s[i] = s[i+1]
    s[i+1] = temp
  }
}