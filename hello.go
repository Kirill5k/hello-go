package main

import (
	"fmt"
	"time"
)

const (
	HelloWorldTemplate string = "Hello, %s!"
)

func main() {
	//var declaration, only works inside functions
	helloStr := fmt.Sprintf(HelloWorldTemplate, "Boris")
	fmt.Println(helloStr)
	fmt.Printf("The type is %T\n", helloStr)

	dates()
	pointers()
}

func dates() {
	now := time.Now()
	fmt.Println("Today's date is " + now.Format(time.RFC3339))
}

func pointers() {
	anInt := 42
	var pointer = &anInt
	fmt.Println("Value of pointer:", *pointer)
}
