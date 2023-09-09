package main

import (
	"fmt"
)

func main() {

	//declare a var with a type
	var helloStr string = "Hello, %s!"
	//var declaration, only works inside functions
	nameStr := "World"

	fmt.Println(fmt.Sprintf(helloStr, nameStr))
	fmt.Printf("The type is %T\n", helloStr)
}
