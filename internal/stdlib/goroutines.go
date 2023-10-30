package stdlib

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func GoDates() {
	now := time.Now()
	fmt.Println("Today's date is " + now.Format(time.RFC3339))
}

func GoPointers() {
	anInt := 42
	var pointer = &anInt
	fmt.Println("Value of pointer", *pointer)
}

func GoArrays() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println("Colors array", colors)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers array", numbers)
}

func GoSlices() {
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

func GoMaps() {
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

func GoStructs() {
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

func GoInterfaces() {
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

func GoGenerics() {
	minFloat, _ := min([]float64{2, 3, 5})
	fmt.Println("Generic ordered interface with float", minFloat)

	minString, _ := min([]string{"B", "A", "C"})
	fmt.Println("Generic ordered interface with string", minString)
}

func GoConditionalLogic() {
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

func GoSwitch() {
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

func GoLoops() {
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

func GoFunctions() {
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
	Status          string `json:"status"`
	StartupTime     string `json:"startup_time"`
	UpTime          string `json:"up_time"`
	ServerIpAddress string `json:"server_ip_address"`
}

func GoRecovery() {
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

func GoSelect() {
	ch1, ch2 := make(chan int), make(chan int)
	go func() {
		ch1 <- 42
	}()

	select {
	case val := <-ch1:
		fmt.Printf("Received value from ch1: %d\n", val)
	case val := <-ch2:
		fmt.Printf("Received value from ch2: %d\n", val)
	}

	chOut := make(chan float64)
	go func() {
		time.Sleep(100 * time.Millisecond)
		chOut <- 3.14
	}()

	select {
	case val := <-chOut:
		fmt.Printf("Received %f\n", val)
	case <-time.After(20 * time.Millisecond):
		fmt.Println("timeout")
	}
}

func GoContext() {

	type Bid struct {
		AdUrl string
		Price float64
	}

	defaultBId := Bid{AdUrl: "https://adsrus.com/default", Price: 0.02}

	bestBid := func(url string) Bid {
		time.Sleep(20 * time.Millisecond)
		return Bid{AdUrl: "https://adsrus.com/19", Price: 0.05}
	}

	findBid := func(ctx context.Context, url string) Bid {
		ch := make(chan Bid, 1) // buffered channel to avoid goroutine leak
		go func() {
			ch <- bestBid(url)
		}()

		select {
		case bid := <-ch:
			return bid
		case <-ctx.Done():
			return defaultBId
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	bid := findBid(ctx, "https://http.cat/418")
	fmt.Printf("Found bid %+v\n", bid)
}

func GoHttpClient() {
	resp, _ := http.Get("https://reqfol.fly.dev/health/status")
	fmt.Printf("Response: %+v\n", resp)
	defer func(Body io.ReadCloser) {
		fmt.Printf("closing resp.Body")
		if err := Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)

	var status HealthStatus
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&status); err != nil {
		panic(err)
	}

	fmt.Printf("HealthStatus: %+v\n", status)
}

func GoHttpContextClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 3000*time.Millisecond)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://reqfol.fly.dev/health/status", nil)
	resp, _ := http.DefaultClient.Do(req)
	fmt.Printf("Response: %+v\n", resp)
	defer func(Body io.ReadCloser) {
		fmt.Printf("closing resp.Body")
		if err := Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)

	var status HealthStatus
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&status); err != nil {
		panic(err)
	}

	fmt.Printf("HealthStatus: %+v\n", status)
}
