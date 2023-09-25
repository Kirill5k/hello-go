package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	HelloWorldTemplate string = "Hello, %s!"
)

/*
https://go-proverbs.github.io for reference
*/
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
	goInterfaces()
	goGenerics()
	goConditionalLogic()
	goSwitch()
	goLoops()
	goFunctions()
	goHttp()
	goRecovery()
	goRoutines()
	goChannels()
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

	_, ok := states["AL"]
	if !ok {
		fmt.Println("AL not found in states map")
	}

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

func NewDog(breed string, age int) (*Dog, error) {
	if age <= 0 {
		return nil, fmt.Errorf("age must be greater than 0")
	}

	dog := Dog{breed, age}
	return &dog, nil
}

func (d Dog) Burk() string {
	return "Woof"
}

func (d *Dog) SetAge(age int) {
	d.Age = age
}

func goStructs() {
	poodle, err := NewDog("Poodle", 10)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	poodle.SetAge(42)
	fmt.Println(poodle, poodle.Burk())
	fmt.Printf("%+v\n", poodle) // debug
	fmt.Printf("%#v\n", poodle) // debug with types
}

type Shape interface {
	Area() float64
}

type Square struct {
	Length float64
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func goInterfaces() {
	sumAreas := func(shapes []Shape) float64 {
		total := 0.0
		for _, shape := range shapes {
			total += shape.Area()
		}
		return total
	}

	shapes := []Shape{Circle{10.5}, Square{2.5}}

	fmt.Println("Total area of shapes is", sumAreas(shapes))
}

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](items []T) (T, error) {
	if len(items) == 0 {
		var zero T
		return zero, fmt.Errorf("min of empty slice")
	}

	m := items[0]
	for _, i := range items[1:] {
		if i < m {
			m = i
		}
	}
	return m, nil
}

func goGenerics() {
	minFloat, _ := min([]float64{2, 3, 5})
	fmt.Println("Generic ordered interface with float", minFloat)

	minString, _ := min([]string{"B", "A", "C"})
	fmt.Println("Generic ordered interface with string", minString)
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

	divmod := func(val1, val2 int) (int, int) {
		return val1 / val2, val1 % val2
	}

	div, mod := divmod(1, 3)

	fmt.Println("Using lambda to add values", sum(1, 3))
	fmt.Println("Using lambda to divide values", div, mod)
}

type HealthStatus struct {
	Status          string
	StartupTime     string `json:"startup_time"`
	UpTime          string `json:"up_time"`
	ServerIpAddress string `json:"server_ip_address"`
}

func goHttp() {
	resp, _ := http.Get("https://reqfol.fly.dev/health/status")
	fmt.Printf("Response: %T\n", resp)
	defer resp.Body.Close()

	var status HealthStatus
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&status)

	fmt.Printf("HealthStatus: %T\n", status)
}

func goRecovery() {
	// Named return values - local variables inside function
	safeValue := func(vals []int, index int) (n int, err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("%v", e)
			}
		}()

		return vals[index], nil
	}

	_, err := safeValue([]int{}, 10)
	fmt.Printf("error from saveValue %v\n", err)
}

func goRoutines() {
	contentType := func(url string) {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		defer resp.Body.Close()
		ctype := resp.Header.Get("content-type")
		fmt.Printf("%s -> %s\n", url, ctype)
	}

	urls := []string{
		"https://golang.com",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			contentType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func goChannels() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	val := <-ch
	fmt.Printf("Meaning of life is %d\n", val)
}
