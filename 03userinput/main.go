package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "wellcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for our pizza:") // auto \n

	// comma ok syntax or error ok syntax

	// input,err :=reader.ReadString('\n')  // we are intersted in input as well as error
	// _ , err  // we are only intersted in err not input 
	input, _ := reader.ReadString('\n')  // not looking at the error
	fmt.Println("thanks for rating,",input)
	fmt.Printf("type of this rating is %T",input) // the type is string
}