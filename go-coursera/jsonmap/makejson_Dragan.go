package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name	string	`json:"name"`
	Address string	`json:"address"`
}

func main() {
	p1 := Person{}
	fmt.Print("Please enter your name: ")
	_, err := fmt.Scan(&p1.Name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("Please enter your address: ")
	_, err = fmt.Scan(&p1.Address)
	if err != nil {
		fmt.Println(err)
	}

	prettyJSON, err := json.MarshalIndent(p1, "", "    ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%s\n", string(prettyJSON))

}