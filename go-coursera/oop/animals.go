package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food);
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion);
}

func (a *Animal) Speak() {
	fmt.Println(a.noise);
}

func main() {
	var strEval string
	var arrTmp []string

	// "infinite" loop to read the user inputs
	for {
		fmt.Printf(`Please enter "animal action" or 'X' to quit> `)
		// Use bufio to accept spaces into the string (i.e: "cow food")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			strEval = scanner.Text()
		}
		// if the user enters 'x'/'X' break/stop the reading
		if strings.ToUpper(strEval) == "X" {
			break
		}
		
		arrTmp = strings.Split(strEval, " ");
		if len(arrTmp) != 2 {
			fmt.Println("Not enough arguments");
			continue
		}
		
		animal, error := GetAnimal(arrTmp[0])
		if error != nil {
			fmt.Println(error);
			continue
		}
		PrintAnimalInfo(animal, arrTmp[1])
	}
}

func GetAnimal(strAnimal string) (*Animal, error) {
		switch(strAnimal) {
			case "cow":
				return &Animal{
						food: "grass",
						locomotion: "walk",
						noise: "moo",
				}, nil
			case "bird":
				return &Animal{
						food: "worms",
						locomotion: "fly",
						noise: "peep",
				}, nil
			case "snake":
				return &Animal{
						food: "mice",
						locomotion: "slither",
						noise: "hsss",
				}, nil
			default:
				return nil, errors.New("Not a valid Animal.")
		}
}

func PrintAnimalInfo(animal *Animal, strAction string) {
	switch(strAction) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Not a valid animal action.");
	}
}