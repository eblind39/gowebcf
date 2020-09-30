package main 

import(
	"fmt"
	"strings"
)

type Animal struct{
	name string
	eat string
	move string
	speak string 
}
func (ani Animal) Eat(){
	fmt.Println(ani.eat)//ani.name+"'s eat "+ani.eat)
}

func (ani Animal) Move(){
	fmt.Println(ani.move)//ani.name+"'s moves by "+ani.move+"ing")
}

func (ani Animal) Speak(){
	fmt.Println(ani.speak)//ani.name+"'s speak "+ani.speak)
}

func main(){
	cow := Animal{name:"cow",
				  eat:"grass",
				  move: "walk",
				  speak: "moo"}

	var bird Animal
		bird.name="bird"
		bird.eat="worn"
		bird.move="fly"
		bird.speak="peep"

	var snake Animal
		snake.name="snake"
		snake.eat="mice"
		snake.move="slither"
		snake.speak="hsss"

	var ani,meth string
	var animal Animal
	for {
    	fmt.Println("Animal \t method()") 
    	fmt.Print(">") 						// > cow move		
    	fmt.Scan(&ani)						// cow moves by walking
    	if (strings.ToLower(ani)=="snake"){
    		animal=snake
    	} else if (strings.ToLower(ani)=="bird"){
    		animal=bird
    	} else {
    		animal=cow
    	}
    
    	fmt.Scan(&meth)
    	if (strings.ToLower(meth[0:1])=="e"){
    		animal.Eat()
    	} else if (strings.ToLower(meth[0:1])=="m"){
    		animal.Move()
    	} else {
    		animal.Speak()
    	}
    	fmt.Println()
    }
}