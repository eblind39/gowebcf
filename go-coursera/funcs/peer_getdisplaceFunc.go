package main

import (
   "bufio"
   "fmt"
   "math"
   "os"
   "strconv"
   "strings"
)

//GenDisplaceFn definition
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
   fn := func(time float64) float64 {
      return 0.5*acceleration*math.Pow(time, 2) + initialVelocity*time + initialDisplacement
   }

   return fn
}

func main() {
   //var acceleration, initialVelocity, initialDisplacement, time float64

   for {
      fmt.Println("please input the acceleration, initialVelocity, initialDisplacement, time separated by space")

      scanner := bufio.NewScanner(os.Stdin)
      scanner.Scan()
      input := scanner.Text()
      input = strings.Trim(input, " ")

      fmt.Printf("input is: [%s]\n", input)

      acceleration, _ := strconv.ParseFloat(strings.Split(input, " ")[0], 64)
      initialVelocity, _ := strconv.ParseFloat(strings.Split(input, " ")[1], 64)
      initialDisplacement, _ := strconv.ParseFloat(strings.Split(input, " ")[2], 64)
      time, _ := strconv.ParseFloat(strings.Split(input, " ")[3], 64)

      fmt.Printf("acceleration: %f\n", acceleration)
      fmt.Printf("initialVelocity: %f\n", initialVelocity)
      fmt.Printf("initialDisplacement: %f\n", initialDisplacement)
      fmt.Printf("time: %f\n", time)

      fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

      //fmt.Println(fn(3))
      fmt.Println(fn(time))
   }
}