package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
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
	goConditionalLogic()
	goSwitch()
	goLoops()
	goFunctions()
	goHttp()
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
	var colors = []string{"Red", "Green", "Blue"}
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

func (d Dog) Burk() string {
	return "Woof"
}

func goStructs() {
	poodle := Dog{"Poodle", 10}

	fmt.Println(poodle, poodle.Burk())
	fmt.Printf("%+v\n", poodle)
}

func goConditionalLogic() {
	var result string

	// Initialises value as part of if statement
	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println("Result is", result)
}

func goSwitch() {
	dow := rand.Intn(7) + 1

	var result string
	switch dow {
	case 1:
		result = "Sunday"
	case 2:
		result = "Monday"
	default:
		result = "Some other day"
	}

	fmt.Println("The day is", result)
}

func goLoops() {
	colors := []string{"Red", "Green", "Blue"}

	for i := 0; i < len(colors); i++ {
		fmt.Print("color ", colors[i], " ")
	}

	fmt.Print("\n")

	for i := range colors {
		fmt.Print("color ", colors[i], " ")
	}

	fmt.Print("\n")

	for _, color := range colors {
		fmt.Print("color ", color, " ")
	}

	fmt.Print("\n")
}

func goFunctions() {
	sum := func(val1, val2 int) int {
		return val1 + val2
	}

	fmt.Println("Using lambda to add values", sum(1, 3))
}

func goHttp() {
	resp, _ := http.Get("https://reqfol.fly.dev/health/status")
	fmt.Println("Response: %T\n", resp)
	defer resp.Body.Close()

	bytes, _ := io.ReadAll(resp.Body)
	content := string(bytes)

	fmt.Println(content)
}
