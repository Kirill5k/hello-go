package main

import (
	"fmt"
)

func main() {

	//declare a var with a type
	var helloStr string = "Hello, "
	//var declaration, only works inside functions
	nameStr := "World"

	fmt.Println(helloStr + nameStr + "!")
	fmt.Printf("The type is %T\n", helloStr)
}
