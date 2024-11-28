package main

import "fmt"

func main() {
	fmt.Println("welcome to class on pointers")
// variable
	// var one int= 2
	// fmt.Println(one)
// pointer
	// var ptr *int
	// fmt.Println("value of printer is", ptr) //value of printer is <nil>

	myNumber := 323

	var  abc = &myNumber  // reference means &  // pinter directly reffer the adderss of the reference
	fmt.Println("value of actual pointer is ",abc) //value of actual pointer is  0xc00000a0e8
	fmt.Println("value of actual pointer is ",myNumber) //value of actual pointer is  323

	*abc = *abc * 2  // operation will be perform on the actural value not the address
	fmt.Println("new value is:", myNumber) // new value is: 646

}
