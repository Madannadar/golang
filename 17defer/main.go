package main

import "fmt"

func main() {
	// fmt.Println("defer in golang")
	// defer fmt.Println("hello world")

	defer fmt.Println("hello world") // defer means it will run last 
	fmt.Println("defer in golang") // output:defer in golang \n hello world

	defer fmt.Println("one")  // last in first out (lifo)
	defer fmt.Println("two")  //three /n two /n one
	defer fmt.Println("three")

// defer in golang ,three, two, one, hello world

	myDefer() //43210

//defer in golang,43210three,two,one,hello world
}

func myDefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
