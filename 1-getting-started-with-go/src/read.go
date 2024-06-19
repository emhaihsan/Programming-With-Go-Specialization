package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main(){
	var filename string
	var names []Name
	fmt.Print("Enter the name of the file: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		name := Name{
			Fname: parts[0],
			Lname: parts[1],
		}
		names = append(names, name)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	for _, name := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n",name.Fname, name.Lname)
	}
}