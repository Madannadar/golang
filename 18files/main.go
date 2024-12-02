package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "this needs to go in a file madan is getting better"

	file, err := os.Create("./mygofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNIlError(err)

	length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }
	checkNIlError(err)
	
	fmt.Println("length is:", length)

	defer file.Close()

	readFile("./mygofile.txt")
}

func readFile(filename string){
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	checkNIlError(err)

	fmt.Println("Text data inside the file is \n", databyte)
}

func checkNIlError(err error)  {
	if err != nil {
		panic(err)
	}
}
