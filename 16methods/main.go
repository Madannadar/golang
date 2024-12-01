// on inheritance in golang; no super or parent
package main

import "fmt"

func main() {
	fmt.Println("structs in golang")

	madan := User{"madan", "test.email", true, 20}
	fmt.Println(madan)                             //{madan test.email true 20}
	fmt.Printf("madan details are : %+v\n", madan) //madan details are : {Name:madan Email:test.email Status:true Age:20}
	fmt.Printf("Name is %v and email is %v.\n", madan.Name, madan.Email)
	madan.GetStatus()
	madan.NewMain()
	fmt.Printf("Name is %v and email is %v.\n", madan.Name, madan.Email)
}


type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// oneAge int  //first letter cap to make it exportable
}

func (u User) GetStatus()  {  // u User means u is the short form for User
	fmt.Println("Is user active:",u.Status)
}

func ( a User) NewMain(){
	a.Email = "test.gmail.com"  // does not changes the main it changes a copy
	fmt.Println("Email of this user is:",a.Email)
}