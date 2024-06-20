package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1] // remove \n

	fmt.Print("Enter an address: ")
	address, _ := reader.ReadString('\n')
	address = address[:len(address)-1] // remove \n

	data := map[string]string{
		"name": name,
		"address": address,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))
}