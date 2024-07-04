// this program aims at reversing a string
package main

import "fmt"

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

func main() {
	//prompting user for a string they'd like revrsed

	var str string
	fmt.Println("What String would you like reversed?: ")
	fmt.Scan(&str)

	//then call the function,by passing user input as its argument
	//while assigning it to another variable reversedStr for easy printing on the screen

	reverseStr := reverseString(str)

	fmt.Printf("\n OriginalString: %v \n Reversed String: %v\n", str, reverseStr)
}
