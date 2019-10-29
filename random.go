package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // without this line, rand will give you the same number everytime
	amountLeft := rand.Intn(10000) 
	fmt.Println("amountLeft is: ", amountLeft)
	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}