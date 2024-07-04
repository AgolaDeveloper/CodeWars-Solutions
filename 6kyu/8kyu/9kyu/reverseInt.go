package main

import (
	"fmt"
	"strconv"
)

//defining a func that receives an integer number and reverses its digits
//... 2315->5132

func reverseInt(number int) int {
	//first convert the INTEGER number into a STRING
	str := strconv.Itoa(number)
	//str := string(number)

	//now it's the string that we need to reverse
	//...but must ensure it's converted to a Slice
	//slice of runes

	ourSlice := []rune(str)

	//then implement an algo to reverse the elements in our slice of runes

	//now, we convert the reversed slice of runes back to a string

	str1 := string(ourSlice)

	//remember we need to again convert the string back to an integer
	//num := int(str1)
	num, _ := strconv.Atoi(str1)

	//now we return our already-converted integer number
	return num

}

func main() {

	var num int
	fmt.Println("Enter your int number that you'd like us to REVERSE: ")
	fmt.Scan(&num)

	reversedNum := reverseInt(num)
	fmt.Printf("\nHere's your revrsed Integer number: %v\n", reversedNum)
}
