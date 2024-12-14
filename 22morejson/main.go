// watch the video again
package main

import (
	"encoding/json"
	"fmt"
)

type course struct { // lowercase  // aliases
	Name     string `json:"coursename"`  // giving a name to display the values
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"` // - means will not show 
	Tags     []string `json:"tags,omitempty"` // slice of string // omit if there is nil in data 
}

func main() {
	fmt.Println("Welcome to More Json")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearncodeOnline.in", "123456", []string{"Web-dev", "js"}},
		{"ReactJs GN", 199, "LearncodeOnline.in", "123d456", []string{"Web-frontend", "js"}},
		{"ReactJs GM", 599, "LearncodeOnline.in", "123s456", []string{"Web-backend", "js"}},
		{"ReactJs GA", 399, "LearncodeOnline.in", "123f456", nil},
	}

	// 	// package this data as Json data

	// finalJson, err := json.Marshal(lcoCourses)
	// Welcome to More Json
	// [{"Name":"ReactJs Bootcamp","Price":299,"Password":"123456","Tags":["Web-dev","js"]},{"Name":"ReactJs GN","Price":199,"Password":"123d456","Tags":["Web-frontend","js"]},{"Name":"ReactJs GM","Price":599,"Password":"123s456","Tags":["Web-backend","js"]},{"Name":"ReactJs GA","Price":399,"Password":"123f456","Tags":null}]

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	// Welcome to More Json
	// [
	//         {
	//                 "Name": "ReactJs Bootcamp",
	//                 "Price": 299,
	//                 "Password": "123456",
	//                 "Tags": [
	//                         "Web-dev",
	//                         "js"
	//                 ]
	//         },
	//         {
	//                 "Name": "ReactJs GN",
	//                 "Price": 199,
	//                 "Password": "123d456",
	//                 "Tags": [
	//                         "Web-frontend",
	//                         "js"
	//                 ]
	//         },
	//         {
	//                 "Name": "ReactJs GM",
	//                 "Price": 599,
	//                 "Password": "123s456",
	//                 "Tags": [
	//                         "Web-backend",
	//                         "js"
	//                 ]
	//         },
	//         {
	//                 "Name": "ReactJs GA",
	//                 "Price": 399,
	//                 "Password": "123f456",
	//                 "Tags": null
	//         }
	// ]
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", finalJson)
}

// consume json data video 32

func DecodeJson(){
	jsonDataFromWb := []byte(`
	{	"coursename": "ReactJs GN",
        "Price": 199,
        "website": "LearncodeOnline.in",
        "tags": ["Web-frontend","js"]
	}
	`) // web data comes in byte forate

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}else{
		fmt.Println("Json was not valid")
	}

	// some case where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWb, &myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k , v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n",k,v,v)
	}
}

