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
	arrays()
	slices()
}

func dates() {
	now := time.Now()
	fmt.Println("Today's date is " + now.Format(time.RFC3339))
}

func pointers() {
	anInt := 42
	var pointer = &anInt
	fmt.Println("Value of pointer", *pointer)
}

func arrays() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Colors array", colors)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array", numbers)
}

func slices() {
	var colors = []string{"Red", "Green", "Bluee"}
	colors = append(colors, "Purple")
	colors = append(colors[1:])
	fmt.Println("Colors slice", colors)

	numbers := make([]int, 5, 5) // last argument is optional slice capacity
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	fmt.Println("Numbers slice", colors)
}
