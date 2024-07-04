// this program aims at reversing a string
package main

import (
	"fmt"
	"strconv"
)

//this function receives a string and returns its Reversed value

func reverseString(str string) string {
	//we first convert the string into a slice
	//...slice of runes in this case

	runes := []rune(str)

	//implement an algorithm that reverses slice elements on runes

	n := len(runes) - 1

	for i := n; i >= n-i; i-- {
		for j := n - i; j <= n-i; j++ {
			//then swap umbers from both end moving towards the center

			temp := runes[i]
			runes[i] = runes[j]
			runes[j] = temp
		}
	}

	//we now have reversed slice of runes
	//it's time to convert it back to a string

	ourString := string(runes)

	//then return the revrsed string

	return ourString
}

//converting STRING to a Integer NUMBER
//use methods from strconv package to help with conversion between STRINGS and INTEGERS

func str2Int(str string) {

	fmt.Println("_______________________________________________________")

	fmt.Println("CONVERTING STRING TO INTEGER")

	//STRING to INTEGER we use Atoi() method
	ourInt, err := strconv.Atoi(str)

	if err != nil {
		fmt.Printf("\nError Converting %v to an Integer\nERROR!: %v\n", str, err)
	} else {
		//now return ourInt; already-converted from string
		fmt.Println("Original string: ", str)
		fmt.Println("Converted Result (INTEGER) string: ", ourInt)
	}

}

func main() {
	//prompting user for a string they'd like revrsed

	var str string
	fmt.Println("What String would you like reversed?: ")
	fmt.Scan(&str)

	//then call the function,by passing user input as its argument
	//while assigning it to another variable reversedStr for easy printing on the screen

	reverseStr := reverseString(str)

	fmt.Printf("\n OriginalString: %v \n Reversed String: %v\n", str, reverseStr)

	concert := "Reggae"
	str2Int(concert)

}
