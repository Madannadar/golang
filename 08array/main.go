// no sorting and array is not much used in golang slices is used

package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to array in golang")

	var fruitList [4]string

	 fruitList[0] = "Apple"
	 fruitList[1] = "orange"
	 fruitList[3] = "grape" // here as 2nd element is empty it output a space

	 fmt.Println("fruit list is:", fruitList) //fruit list is: [Apple orange  grape]
	 fmt.Println("fruit list is:", len(fruitList)) //fruit list is: 4

	 var vegList = [3]string{"potato","beans","mushroom"}
	 fmt.Println("vegy list is:",vegList)
	 fmt.Println("vegy list is:",len(vegList))
}

//-------------------------slcies--------------------------------
