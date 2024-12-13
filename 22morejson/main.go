// watch the video again
package main

import "fmt"

type course struct{
	Name string
	Price int
	platform string
	Password string
	Tags 	[]string
}

func main() {
	fmt.Println("Welcome to More Json")
	
}


// func EncodeJson(){

// 	lcoCourses := []course{
// 		{"ReactJs Bootcamp",299,"LearncodeOnline.in","123456",[]string{"Web-dev","js"}},
// 		{"ReactJs GN",299,"LearncodeOnline.in","123456",[]string{"Web-frontend","js"}},
// 		{"ReactJs GM",299,"LearncodeOnline.in","123456",[]string{"Web-backend","js"}},
// 		{"ReactJs GM",299,"LearncodeOnline.in","123456",nil},
// 	},
	
// 	// package this data as Json data
// 	finalJson := json.Marshal(lcoCourse)