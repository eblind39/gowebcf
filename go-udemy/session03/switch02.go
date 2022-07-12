package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	
	switch {
		case 2 == 4 || 9 == 10:
			fmt.Println("No debería imprimir 2==4")
		case 2 == 2:
			fmt.Println("Debería imprimir 2==2")
		case 2 == 5 && 3 == 1:
			fmt.Println("No debería imprimir 2==5")
		default:
			fmt.Println("Imprimiendo desde default")
	}
	
	switch i:=43; i {
		case 42, 43:
			fmt.Println("Debería imprimir i==42 || i==43")
		default:
			fmt.Println("Imprimiendo desde default")
	}
	
	switch deporteFav:= "surfing"; deporteFav {
		case "baseball":
			fmt.Printf("%s - ve al estadio", deporteFav)
		case "natación":
			fmt.Printf("%s - ve a la piscina", deporteFav)
		case "surfing":
			fmt.Printf("%s - ve a la playa", deporteFav)
	}
}