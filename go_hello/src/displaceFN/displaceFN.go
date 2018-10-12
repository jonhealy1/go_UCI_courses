package main

import (
  "fmt"
)

func main() {
  var velocity float64
  var displacement float64
  var acceleration float64
  var time float64
  
  fmt.Printf("Please enter a value for the initial displacement:\n")
  fmt.Scan(&displacement)

  fmt.Printf("Please enter a value for the initial velocity:\n")
  fmt.Scan(&velocity)

  fmt.Printf("Please enter a value for the acceleration:\n")
  fmt.Scan(&acceleration)

  fmt.Printf("At what time you would like to know your new displacement at?:\n")
  fmt.Scan(&time)

  fn := GenDisplaceFn(acceleration , displacement, velocity)

  //fmt.Printf("Your new displacement is:\n")
  fmt.Println("Your new displacement at time t =",time, "is:",fn(time))
}

func GenDisplaceFn (a, d, v float64) func (float64) float64 {
      fn:= func (t float64) float64 {
        return (0.5 * a * t * t) + (v * t) + d
      }
    return fn
}
/*
func GenDisplaceFn (a, d, v float64) func (float64) float64 {
      return func(t float64) float64 {
        return a+d+v+t
      }
}
*/