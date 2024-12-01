// func name(parameter1, parameter2) what kind of value to return(signatures){
	// return parameter1 + parameter 2}
package main

import "fmt"
// main function is like int main in c language
// cannot write function inside the function 
func main() {  // main is the first function that will run
	fmt.Println("functions on golang")
	greeter()
//error
	// func greeterTwo()  {
	// 	fmt.Println("another method")
	// }
	greeterTwo() // can call as many function we want
	result := adder(3,5)
	fmt.Println("Result is:", result)

	proResult := proAdder(2,3,4,5,6,67,4)
	fmt.Println("pro result is :",proResult)
}

func adder(valOne int, valTwo int) int{  // the signature can be more the none like (int, string)
										// then we have to give input and output it like we did in result parameter
 	return valOne + valTwo
}

func proAdder(values ... int) int {  // no limit of values
	total := 0
	for _,val := range values{
		total += val
	}
	return total
}

func greeterTwo()  {
	fmt.Println("another method")
}

func greeter()  {
	fmt.Println("Namstey from golang")
}