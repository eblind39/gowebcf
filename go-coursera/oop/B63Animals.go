/*

Write a program which allows the user to get information about a predefined set of animals. 
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. 
The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 
2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the 
three animals and their associated data which should be hard-coded into your program.

Animal		Food eaten	Locomotion method	Spoken sound
cow			grass		walk				moo
bird		worms		fly					peep
snake		mice		slither				hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. 
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. 
Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. 
The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the 
information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by 
printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct 
containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), 
and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s 
food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. 
Your program should call the appropriate method when the user makes a request.

*/

package main 
  
import ( 
    "fmt"
) 

type Animal struct
{
	food string
	locomotion string
	sound string
}

func (p Animal) Eat() {
	fmt.Println(p.food)
}

func (q Animal) Move() {
	fmt.Println(q.locomotion)
}

func (r Animal) Speak() {
	fmt.Println(r.sound)
}

// Main function 
func main() { 
      
	// Get input num from user via keyboard  
	fmt.Println("Information of Cow")
	var c_food, c_move, c_sound string
	fmt.Println("Food eaten:")
	fmt.Scanln(&c_food)
	fmt.Println("Locomotion method:")
	fmt.Scanln(&c_move)
	fmt.Println("Spoken sound:")
	fmt.Scanln(&c_sound)
	
	fmt.Println("Information of Bird")
	var b_food, b_move, b_sound string
	fmt.Println("Food eaten:")
	fmt.Scanln(&b_food)
	fmt.Println("Locomotion method:")
	fmt.Scanln(&b_move)
	fmt.Println("Spoken sound:")  
	fmt.Scanln(&b_sound) 
	
	fmt.Println("Information of Snake")
	var s_food, s_move, s_sound string
	fmt.Println("Food eaten:")
	fmt.Scanln(&s_food)
	fmt.Println("Locomotion method:")
	fmt.Scanln(&s_move)
	fmt.Println("Spoken sound:") 
	fmt.Scanln(&s_sound)  
	
	var animal, feature string
	for{
		
		fmt.Println("> Animal Name [cow/bird/snake]")
		fmt.Scanln(&animal)
		fmt.Println("> Feature [eat / move / speak] :")
		fmt.Scanln(&feature)

		if(animal == "cow" || animal == "Cow"){
			
			var cow Animal= Animal{c_food,c_move,c_sound}

			if(feature == "eat"){
				cow.Eat()
			}
			if(feature == "move"){
				cow.Move()
			}
			if(feature == "speak"){
				cow.Speak()
			}
		}
		
		if(animal == "bird" || animal == "Bird"){

			var bird Animal= Animal{b_food, b_move, b_sound}

			if(feature == "eat"){
				bird.Eat()
			}
			if(feature == "move"){
				bird.Move()
			}
			if(feature == "speak"){
				bird.Speak()
			}
		}

		if(animal == "snake" || animal == "Snake"){

			var snake Animal= Animal{s_food, s_move, s_sound}

			if(feature == "eat"){
				snake.Eat()
			}
			if(feature == "move"){
				snake.Move()
			}
			if(feature == "speak"){
				snake.Speak()
			}
		}		
		  	 
	}	     
}	   