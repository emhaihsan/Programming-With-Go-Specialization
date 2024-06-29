package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
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


func main(){
	animals := map[string] Animal{
		"cow": {"grass","walk","moo"},
		"bird": {"worms", "fly", "peep"},
		"snake" : {"mice", "slither", "hsss"},
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		if strings.ToLower(input) == "exit\n" {
			break
		}
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")

		if len(parts) != 2 {
			fmt.Println("Invalid input, please enter an animal and action '<animal><action>'")
			continue
		}

		animalName := strings.ToLower(parts[0])
		action := strings.ToLower(parts[1])

		animal, ok := animals[animalName]
		if !ok {
			fmt.Println("Invalid animal, please enter a valid animal")
			continue
		}

		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Invalid action, please enter a valid action")
		}
	}
		
}
