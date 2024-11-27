// gives error if a variable is declared but not used
// uint8 = 0 to 255
// uint16 = 0 to 65535
// uint64 = 0 to 18444432 very long 615 last 3 digit
package main

import "fmt"

const LoginToken string = "gfsfsdf" // the first leter L is cap so it is 'public' constent can be used in any file in the folder

func main() {
	var username string = "madan"
	fmt.Println(username)
	fmt.Printf("var is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("var is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 // no error
	// var smallVal uint8 = 256 // yes error
	fmt.Println(smallVal)
	fmt.Printf("var is of type: %T \n", smallVal)

	var smallFloat float32 = 25.434324234 // get 5 value after .
	// var smallFloat float64 = 25.545323523453 // give all value after .
	fmt.Println(smallFloat)
	fmt.Printf("var is of type: %T \n", smallFloat)

	var anothervariable int // ans 0
	// var anothervariable string // ans blank line
	fmt.Println(anothervariable)
	fmt.Printf("var is of type: %T \n", anothervariable)

	// implicit type

	var website = "google.com" // here we have not specified the type of var so lexer will decide with the value and then only that type of value will only be accepted for the variable
	fmt.Println(website)
	// website = 23// error

	numberofUser := 3000 //most common syntax but only used inside a method or function not outside
	fmt.Println(numberofUser)

	fmt.Println(LoginToken) // calling var that is public out of the function
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
