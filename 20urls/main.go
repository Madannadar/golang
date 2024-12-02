package main

import (
	"fmt"
	"net/url"
)

const url1 string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dfdasdf"

func main() {
	fmt.Println("url in golang")
	fmt.Println(url1)

	// parsing
	result, _ := url.Parse(url1)

	// fmt.Println(result.Scheme)  //https
	// fmt.Println(result.Host) //lco.dev:3000
	// fmt.Println(result.Path) //learn
	// fmt.Println(result.Port()) //3000
	// fmt.Println(result.RawQuery) //coursename=reactjs$paymentid=dfdasdf

	qparams := result.Query()
	fmt.Printf("the type of query parms are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("param is:",val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/learn",
		RawQuery: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
