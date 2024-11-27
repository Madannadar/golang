package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of golang")

	presentTime := time.Now()  // givers us the current time 2024-11-27 22:24:16.824231 +0530 IST m=+0.000649201
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006")) // use this time for formating ans: 11-27-2024
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //ans: 11-27-2024 Wednesday
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // ans:11-27-2024 22:27:51 Wednesday

	createdDate :=time.Date(2020, time.August, 1, 5, 5, 5, 5, time.UTC)//hour min sec milisec
	fmt.Println(createdDate) 
	fmt.Println(createdDate.Format("01-02-2006 Monday"))  // monday wrong Monday correct
}
// terminal
//GOOS="darwin" go build 
// building the app for mac 
// GOOS="windows" go build for windows
// GOOS="linux" go build

