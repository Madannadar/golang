package main

import "fmt"

func main() {
	fmt.Println("loops in golang")

	// days := []string{"Sunday","Monday","Tuesday","wednesday","Thursday","Friday","Saturday"}
	// fmt.Println(days)

	// for i:=0; i< len(days); i++{
	// 	fmt.Printf("index:%v day:%v \n",i,days[i])
	// }

	// for i := range days{
	// 	fmt.Printf("index:%v day:%v \n",i,days[i])
	// }

	// for i, day := range days{
	// 	fmt.Printf("index is %v and value is %v \n",i,day)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 6 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Printf("value is : %v \n",rougueValue)
		rougueValue++
	}

	lco:
		fmt.Println("jumping")
}
