package main

import (
	"fmt"
)

// var myName string

func main() {
	fmt.Println("Hello, world!")

	var whatToSay string

	var i int

	whatToSay = "Goodby, cruel world!"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("The function returned", whatWasSaid)
	fmt.Println("The other thing that was said is", theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "Something", "Something else"
}
