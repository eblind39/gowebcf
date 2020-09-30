package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"errors"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}
func (c Cow) Eat() { fmt.Println("grass"); }
func (c Cow) Move() { fmt.Println("walk"); }
func (c Cow) Speak() { fmt.Println("moo"); }

type Bird struct {
	name string
}
func (b Bird) Eat() { fmt.Println("worms"); }
func (b Bird) Move() { fmt.Println("fly"); }
func (b Bird) Speak() { fmt.Println("peep"); }

type Snake struct {
	name string
}
func (s Snake) Eat() { fmt.Println("mice"); }
func (s Snake) Move() { fmt.Println("slither"); }
func (s Snake) Speak() { fmt.Println("hsss"); }

func GetAnimalName(a Animal) string {
	cow, ok := a.(Cow)
	if ok { return cow.name }
	bird, ok := a.(Bird)
	if ok { return bird.name }
	snake, ok := a.(Snake)
	if ok { return snake.name }
	return ""
}

func main() {
	var strEval string
	var arrTmp []string
	var animals []Animal

	fmt.Println(`---Please enter---`)
	fmt.Println(`"newanimal animalname animaltype"`)
	fmt.Println(`"query animalname animalaction"`)
	fmt.Println(`'X' to quit`)
	// "infinite" loop to read the user inputs
	for {
		fmt.Print("\n> ")
		// Use bufio to accept spaces into the string (i.e: "cow food")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() { strEval = scanner.Text() }
		// if the user enters 'x'/'X' break/stop the reading
		if strings.ToUpper(strEval) == "X" { break }
		
		arrTmp = strings.Split(strEval, " ");
		if len(arrTmp) != 3 {
			fmt.Println("Not enough arguments");
			continue
		}
		
		if arrTmp[0] == "newanimal" || arrTmp[0] == "query" {
			if arrTmp[0] == "newanimal" {
				animal, error := GetAnimal(arrTmp[2], arrTmp[1])
				if error != nil {
					fmt.Println(error);
					continue
				}
				animals = append(animals, animal)
				fmt.Println("Created it!");
			}
			if arrTmp[0] == "query" {
				animal, error := GetAnimalByName(animals, arrTmp[1])
				if error != nil {
					fmt.Println(error);
					continue
				}
				PrintAnimalInfo(animal, arrTmp[2])
			}
		} else {
			fmt.Println("Not a valid input");
		}
	}
}

func GetAnimalByName(animals []Animal, strName string) (Animal, error) {
	var strTmp string
	
	for _, animal := range animals {
		switch anml := animal.(type) {
			case Cow, Bird, Snake:
				strTmp = GetAnimalName(anml)
				if strTmp == strName {
					return animal, nil;
				}
		}
	}
	return nil, errors.New("No animal was found with that name.")
}

func GetAnimal(strAnimal string, strName string) (Animal, error) {
		switch(strAnimal) {
			case "cow":
				return Cow{
						name: strName,
				}, nil
			case "bird":
				return Bird{
						name: strName,
				}, nil
			case "snake":
				return Snake{
						name: strName,
				}, nil
			default:
				return nil, errors.New("Not a valid Animal.")
		}
}

func PrintAnimalInfo(animal Animal, strAction string) {
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