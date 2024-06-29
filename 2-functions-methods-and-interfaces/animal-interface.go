package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface{
	Eat()
	Move()
	Speak()
}

type Cow struct{
	name string
}

type Bird struct{
	name string
}

type Snake struct{
	name string
}

func (c Cow) Eat(){
	fmt.Println("grass")
}

func (c Cow) Move(){
	fmt.Println("walk")
}

func (c Cow) Speak(){
	fmt.Println("moo")
}

func (b Bird) Eat(){
	fmt.Println("worms")
}
func (b Bird) Move(){
	fmt.Println("fly")
}

func (b Bird) Speak(){
	fmt.Println("peep")
}
func (s Snake) Eat(){
	fmt.Println("mice")
}

func (s Snake) Move(){
	fmt.Println("slither")
}
func (s Snake) Speak(){
	fmt.Println("hsss")
}

var animals = make(map[string] Animal)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) < 3 {
			fmt.Println("Invalid command")
		}

		command := parts[0]
		animalName := parts[1]
		action := parts[2]

		switch command {
		case "newanimal":
			var animal Animal
			switch action {
			case "cow":
				animal = Cow{name: animalName}
			case "bird":
				animal = Bird{name: animalName}
			case "snake":
				animal = Snake{name: animalName}
			}
			animals[animalName] = animal
		case "query":
			animal, ok := animals[animalName]
			if !ok {
				fmt.Println("Invalid animal")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	}
}