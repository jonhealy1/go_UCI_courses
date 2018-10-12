package main

import (
  "fmt"
  "strconv"
  "sort"
)

func main() {
  var numeral string
  var loop bool
  loop = true
  slice1 := make([]int, 0, 3)
  
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
      slice1 = append(slice1, i1)
      sort.Ints(slice1)
      fmt.Println("slice1", slice1)
    }
  }
}