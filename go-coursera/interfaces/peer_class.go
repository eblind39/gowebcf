package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	argCount      int    = 3
	commandPrompt string = "> "
	quitCode      string = "X"
)

var animalMap map[string]Animal
var animalArray = []string{"cow", "bird", "snake"}
var requestArray = []string{"eat", "move", "speak"}
var commandArray = []string{"newanimal", "query", "show"}

func usage() {
	fmt.Println("Usage 1: newanimal <animal_name> <animal_type>")
	fmt.Println("         where animal_name is a string")
	fmt.Printf("         where animal_type is one of (%s)\n\n", strings.Join(animalArray, ","))
	fmt.Println("Usage 2: query <animal_name> <request>")
	fmt.Println("         where animal_name is a string")
	fmt.Printf("         where request is one of (%s)\n\n", strings.Join(requestArray, ","))
	fmt.Println("Usage 3: show ")
	fmt.Println("         Lists all the animals entered")
	fmt.Println("Enter X to exit...")
}

func init() {
	// https://golang.org/doc/effective_go.html#allocation_new
	animalMap = make(map[string]Animal)
}

func show() {
	for k, v := range animalMap {
		fmt.Printf("%s is a ", k)
		switch v.(type) {
		case Cow:
			fmt.Print("Cow ")
		case Bird:
			fmt.Print("Bird ")
		case Snake:
			fmt.Print("Snake ")
		}
		fmt.Println(v)
	}
}

func createNewAnimal(animalName, animalType string) {
	switch animalType {
	case animalArray[0]:
		animal := Cow{food: "grass", locomotion: "walk", noise: "moo"}
		animalMap[animalName] = animal
	case animalArray[1]:
		animal := Bird{food: "worms", locomotion: "fly", noise: "peep"}
		animalMap[animalName] = animal
	case animalArray[2]:
		animal := Snake{food: "mice", locomotion: "slither", noise: "hsss"}
		animalMap[animalName] = animal
	default:
		fmt.Errorf("%s: Invalid animal type", animalType)
		usage()
		return
	}
	fmt.Println("Created it!")
}

func queryAnimal(animalName, info string) {
	animal := animalMap[animalName]
	switch info {
	case requestArray[0]:
		fmt.Printf("%s eats %s\n", animalName, animal.Eat())
	case requestArray[1]:
		fmt.Printf("%s moves by %s\n", animalName, animal.Move())
	case requestArray[2]:
		fmt.Printf("%s speaks %s\n", animalName, animal.Speak())
	default:
		fmt.Errorf("%s: Invalid Request", info)
		usage()
	}
}

func parseCommand(list []string) {
	if len(list) != argCount {
		fmt.Errorf("Usage error: require %d arguments", argCount)
	}
	switch list[0] {
	case commandArray[0]:
		if itemIsInArray(list[2], animalArray) {
			createNewAnimal(list[1], list[2])
		}
	case commandArray[1]:
		if itemIsInArray(list[2], requestArray) {
			queryAnimal(list[1], list[2])
		}
	case commandArray[2]:
		show()
	}
	fmt.Print(commandPrompt)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	usage()
	fmt.Print(commandPrompt)
	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == quitCode {
			os.Exit(0)
		}
		strVal = stripSpaces(strVal)
		strVal = strings.ToLower(strVal)
		var list = strings.Split(strVal, " ")
		if !(len(list) == argCount && itemIsInArray(list[0], commandArray)) {
			if list[0] != commandArray[2] {
				usage()
				fmt.Print(commandPrompt)
				continue
			}
		}
		parseCommand(list)
	}
}

// Animal Interface
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{ food, locomotion, noise string }

func (a Cow) Eat() string {
	return a.food
}
func (a Cow) Move() string {
	return a.locomotion
}
func (a Cow) Speak() string {
	return a.noise
}

type Bird struct{ food, locomotion, noise string }

func (a Bird) Eat() string {
	return a.food
}

func (a Bird) Move() string {
	return a.locomotion
}

func (a Bird) Speak() string {
	return a.noise
}

type Snake struct{ food, locomotion, noise string }

func (a Snake) Eat() string {
	return a.food
}

func (a Snake) Move() string {
	return a.locomotion
}

func (a Snake) Speak() string {
	return a.noise
}

// Utility Methods
func stripSpaces(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := re_leadclose_whtsp.ReplaceAllString(input, "")
	result = re_inside_whtsp.ReplaceAllString(result, " ")
	return result
}

func itemIsInArray(item string, array []string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}
