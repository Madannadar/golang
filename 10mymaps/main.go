package main

import (
	"fmt"
)

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)// key and value

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["Rb"] = "ruby"

	fmt.Println("list of all languages", languages)
	fmt.Println("Js", languages["js"]) // use key to target a value

	delete(languages,"Rb")
	fmt.Println("list of all languages", languages)

	// loops are intersting in golang 
	// little loop for map

	for key, value := range languages{
		fmt.Printf("For key '%v' , value is '%v'\n",key,value)
	}
}
