package main

import (
  "fmt"
 // "os"
)

type Animal struct {
    food string
    locomotion string
    noise string
}

func (a Animal) Eat() {
  fmt.Println(a.food)
}

func (a Animal) Move() {
  fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
  fmt.Println(a.noise)
}

func main() {

  //var first string
  var cmd1 string
  var cmd2 string
  

for {
  fmt.Println("Please enter an animal (cow, bird, or snake) followed by a space and then an action (eat, move, or speak)")
  fmt.Print(">")
  //cmd1 = os.Args[0]
  //cmd2 = os.Args[1]
  fmt.Scan(&cmd1)
  fmt.Scan(&cmd2)
  
  

  if cmd1 == "cow" {
    cow(cmd2)
  } else if cmd1 == "bird" {
    bird(cmd2)
  } else {
    snake(cmd2)
  }
}
  //fmt.Println(cmd1, cmd2)
}
func cow(s string) {
  cow := Animal {
    food: "grass",
    locomotion: "walk",
    noise: "moo",
  }
  if s == "eat" {
    cow.Eat()
  } else if s == "move" {
    cow.Move() 
  } else {
    cow.Speak()
  }
}
func bird(s string) {
  bird := Animal {
    food: "worms",
    locomotion: "fly",
    noise: "peep",
  }
  if s == "eat" {
    bird.Eat()
  } else if s == "move" {
    bird.Move() 
  } else {
    bird.Speak()
  }
}
func snake(s string) {
  snake := Animal {
    food: "mice",
    locomotion: "slither",
    noise: "hiss",
  }
  if s == "eat" {
    snake.Eat()
  } else if s == "move" {
    snake.Move() 
  } else {
    snake.Speak()
  }
}