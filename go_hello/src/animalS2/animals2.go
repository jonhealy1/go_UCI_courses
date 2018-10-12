package main

import (
  "fmt"
)
type Animal interface {
  Eat() string
  Move() string
  Speak() string
  Name() string
}

type Bird struct {
  name string
}
func (b *Bird) Eat() string {
  return "worms"
}
func (b *Bird) Move() string {
  return "fly"
}
func (b *Bird) Speak() string {
  return "chirp"
}
func (b *Bird) Name() string {
  return b.name
}

type Cow struct {
  name string
}
func (c *Cow) Eat() string {
  return "grass"
}
func (c *Cow) Move() string {
  return "walk"
}
func (c *Cow) Speak() string {
  return "moo"
}
func (c *Cow) Name() string {
  return c.name
}

type Snake struct {
  name string
}
func (s *Snake) Eat() string {
  return "mice"
}
func (s *Snake) Move() string {
  return "slither"
}
func (s *Snake) Speak() string {
  return "hiss"
}
func (s *Snake) Name() string {
  return s.name
}

func main() {
  var cmd1 string
  var cmd2 string
  var cmd3 string
  animals := []Animal{}

  for {
  fmt.Println("1. enter command newanimal, then a space then a name, another space and then the type - cow, bird, or snake")
  fmt.Println("2. enter command query, then a space then type, another space and then the action - eat, move, or speak")
  fmt.Print(">")
  fmt.Scan(&cmd1)
  fmt.Scan(&cmd2)
  fmt.Scan(&cmd3)
  namer := cmd2
 
  if cmd1 == "newanimal" {
    if cmd3 == "bird" {
      var cmd2 Animal
      cmd2 = &Bird{name: namer}
      //fmt.Println(cmd2.Name())
      animals = append(animals,cmd2)
     // fmt.Println(animals)
      fmt.Println("Created")
    } else if cmd3 == "cow" {
      var cmd2 Animal
      cmd2 = &Cow{name: namer}
      animals = append(animals,cmd2)
      fmt.Println("Created")
    } else {
      var cmd2 Animal
      cmd2 = &Snake{name: namer}
      animals = append(animals,cmd2)
      fmt.Println("Created")
    }
  } else if cmd1 == "query" {
    for j := range animals {
      if animals[j].Name() == cmd2 {
        fmt.Println(animals[j].Name())
        if cmd3 == "eat" {
          fmt.Println(animals[j].Eat())
        } else if cmd3 == "speak" {
          fmt.Println(animals[j].Speak())
        } else if cmd3 == "move" {
          fmt.Println(animals[j].Move())
        }
       
      }
    }
  }
  }
}