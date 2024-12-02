package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("lco web request in golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("response type is: %T \n", response) //response type is: *http.Response
	// fmt.Println(response)

	defer response.Body.Close() // caller's responsiblity to close the connection

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databyte) //response type is: *http.Response  to string
	fmt.Println(content)
}
