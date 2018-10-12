package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ print func(string) }

func (c *Cow) Eat()   { print("grass") }
func (c *Cow) Move()  { print("walk") }
func (c *Cow) Speak() { print("moo") }

type Bird struct{ print func(string) }

func (c *Bird) Eat()   { print("worms") }
func (c *Bird) Move()  { print("fly") }
func (c *Bird) Speak() { print("peep") }

type Snake struct{ print func(string) }

func (c *Snake) Eat()   { print("mice") }
func (c *Snake) Move()  { print("slither") }
func (c *Snake) Speak() { print("hsss") }

func main() {
	scaner := bufio.NewScanner(os.Stdin)

	namesToAnimals := map[string]Animal{}
	for {
		fmt.Println()
		fmt.Print(">")
		if !scaner.Scan() {
			continue
		}

		input := scaner.Text()
		terms := strings.Split(input, " ")
		if len(terms) != 3 {
			fmt.Println("Command is unsupported.")
			continue
		}

		command := terms[0]
		switch command {
		case "newanimal":
			newAnimal := createNewAnimal(terms[2])
			if newAnimal != nil{
				namesToAnimals[terms[1]] = newAnimal
				fmt.Println("Created it!")
			}
		case "query":
			animal, ok := namesToAnimals[terms[1]]
			if ok {
				Do(animal, terms[2])
				fmt.Println()
			}
		default:
			fmt.Println("Command is unsupported.")
		}
	}

}

func Do(a Animal, action string) {
	switch action {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func createNewAnimal(animaltype string) Animal {
	switch animaltype {
	case "cow":
		return &Cow{printFmt}
	case "bird":
		return &Bird{printFmt}
	case "snake":
		return &Snake{printFmt}
	}
	return nil
}

func printFmt(s string) {
	fmt.Println(s)
}