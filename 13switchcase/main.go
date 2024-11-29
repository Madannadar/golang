package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("switch case in golang")

	// rand.Seed()(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is", diceNumber)

	switch diceNumber {
	case 1:   
		fmt.Println("Dice value is one")  // no need to add break 
	case 2:
		fmt.Println("Dice value is two") //value of dice is 2
										 //Dice value is two
										 //Dice value is three
		fallthrough
	case 3:
		fmt.Println("Dice value is three")
	case 4:
		fmt.Println("Dice value is four")
	case 5:
		fmt.Println("Dice value is five")
	case 6:
		fmt.Println("Dice value is six")
	default:
		fmt.Println("what what the f**k")
	}
}
