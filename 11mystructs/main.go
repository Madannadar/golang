// on inheritance in golang; no super or parent
package main

import "fmt"

func main() {
	fmt.Println("structs in golang")

	madan := User{"madan","test.email",true,20}
	fmt.Println(madan) //{madan test.email true 20}
	fmt.Printf("madan details are : %+v\n",madan) //madan details are : {Name:madan Email:test.email Status:true Age:20}
	fmt.Printf("Name is %v and email is %v.",madan.Name,madan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

