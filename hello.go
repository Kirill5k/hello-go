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

	goDates()
	goPointers()
	goArrays()
	goSlices()
	goMaps()
	goStructs()
}

func goDates() {
	now := time.Now()
	fmt.Println("Today's date is " + now.Format(time.RFC3339))
}

func goPointers() {
	anInt := 42
	var pointer = &anInt
	fmt.Println("Value of pointer", *pointer)
}

func goArrays() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Colors array", colors)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array", numbers)
}

func goSlices() {
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

func goMaps() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"

	delete(states, "OR")
	states["NY"] = "New York"
	fmt.Println("States map", states)

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}
}

type Dog struct {
	Breed string
	Age   int
}

func goStructs() {
	poodle := Dog{"Poodle", 10}

	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
}
