package main // package declaration
// The package declaration informs the compiler whether to create an executable or library
// a library doesn’t outright run/execute code — it is a collection of code that can be used in other programs
// Programs that have the package declaration package main (like this one) will create an executable

import "fmt" // println is from this package

func main() {
	fmt.Println("Hello World")
}