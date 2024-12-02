// to send the data the request should be same url as the server who gets it in the index.js file and the thunder wala to also send request from the same url in this case http://localhost:8000/postform
package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("get request in golang")
	// performGetRequest()
	// performPostJsonRequst()
	performPostFormRequest()
}
func performGetRequest() {
	const myurl = "http://localhost:8000"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("content length is:", response.ContentLength)

	// method 1
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(content)
	// fmt.Println(string(content))

	//method 2
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("bytecount is:", byteCount)
	fmt.Println(responseString.String())
}

func performPostJsonRequst() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	resquestBody := strings.NewReader(`
		{
			"coursname":"golang",
			"price":0,
			"platform":"learnCodeOnline.in"
		}	
	`)
	response, err := http.Post(myurl, "application/json", resquestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

func performPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data

	data := url.Values{}
	data.Add("firstname", "arun")
	data.Add("lastname", "kumar")
	data.Add("age", "12")
	data.Add("email", "arungandu@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
