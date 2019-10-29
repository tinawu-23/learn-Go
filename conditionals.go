package main

import "fmt"

func main() {
	
	// if else 

	position := 2

	if position == 1 {
	fmt.Println("You won the gold!")
	} else if position == 2 {
	fmt.Println("You got the silver medal.")
	} else if position == 3 {
	fmt.Println("Great job on bronze.")
	} else {
	fmt.Println("Sorry, better luck next time?")
	}

	// switch

	clothingChoice := "sweater"

	switch clothingChoice {
		case "shirt":
			fmt.Println("We have shirts in S and M only.")
		case "polos":
			fmt.Println("We have polos in M, L, and XL.")
		case "sweater":
			fmt.Println("We have sweaters in S, M, L, and XL.")
		case "jackets":
			fmt.Println("We have jackets in all sizes.")
		default:
			fmt.Println("Sorry, we don't carry that")
		}
	}

		// scoped short declaration statements
		x := 8
		y := 9
		if product := x * y; product > 60 {
			fmt.Println(product, "  is greater than 60")
		}
		// fmt.Println(product) <= this will give you undefined product error
		
		switch season := "summer" ; season {
		case "summer":
			fmt.Println("Go out and enjoy the sun!")
		}
	}