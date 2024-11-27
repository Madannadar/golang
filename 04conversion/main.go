// here we are converting the input from string to float but in the input there is a \n automatically added by that l bot so how to convert that to float
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("please rate our pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input , _ := reader.ReadString('\n')
	fmt.Println("thanks for rating,",input)
	// numRating, err := strconv.ParseFloat(input, 64) // converting string to float64  but error one 
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // through this code we have converted the string with \n into float

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("adding 1 to your rating:",numRating+1)
	}
}	