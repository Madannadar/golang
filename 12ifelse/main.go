package main

import "fmt"

func main() {
	fmt.Println("if else in golang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "something else"
	}else{
		result = "exactly 10 login count"
	}
	fmt.Println(result)



	var number int = 20

	if number%2 == 0{
		fmt.Println("number is even")
	}else{
		fmt.Println("number is odd")
	}


	if num :=-1; num < 10{
		fmt.Println("num is less then 10")
	}else{
		fmt.Println("num is not less then 10")
	}

	// if err != nil{
	// }
}
