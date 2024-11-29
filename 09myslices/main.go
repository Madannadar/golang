package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices video")

	var fruitList = []string{"Apple", "Tomato", "Mango"} // no fixed size
	fmt.Printf("type of fruitliset is %T\n", fruitList)  // %T then Printf

	fruitList = append(fruitList, "Banana", "Goa") // add new value into the slice
	// fmt.Println(fruitList)

	fruitList = append(fruitList[1:2]) // the first element is no longer avail [start:end] so 0th position apple is not shown
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 222
	highScores[1] = 221
	highScores[2] = 223
	highScores[3] = 263

	highScores = append(highScores, 444, 333, 111)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove value from slices baseds on index

	var course = []string{"react","javascript","swift","python","ruby"}
	fmt.Println(course)
	var index int = 2;
	course = append(course[:index], course[index+1:]...)  // ... to make the append be used more then once
	fmt.Println(course)
}	

