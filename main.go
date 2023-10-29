package main

import (
	"fmt"
	"kirill5k/hello/go/internal/stdlib"
)

const (
	HelloWorldTemplate string = "Hello, %s!"
)

/*
For reference:
https://go-proverbs.github.io
https://go.dev/doc/effective_go
*/
func main() {
	//var declaration, only works inside functions
	helloStr := fmt.Sprintf(HelloWorldTemplate, "Boris")
	fmt.Println(helloStr)
	fmt.Printf("The type is %T\n", helloStr)

	stdlib.GoDates()
	stdlib.GoPointers()
	stdlib.GoArrays()
	stdlib.GoSlices()
	stdlib.GoMaps()
	stdlib.GoStructs()
	stdlib.GoInterfaces()
	stdlib.GoGenerics()
	stdlib.GoConditionalLogic()
	stdlib.GoSwitch()
	stdlib.GoLoops()
	stdlib.GoFunctions()
	stdlib.GoRecovery()
	stdlib.GoRoutines()
	stdlib.GoChannels()
	stdlib.GoSelect()
	stdlib.GoContext()
	stdlib.GoHttpClient()
	stdlib.GoHttpContextClient()
}
