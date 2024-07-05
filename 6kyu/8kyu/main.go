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

//converting NUMBER to a STRING

func int2Str(num int) {
	//use Itoa()method from strconv package...
	//...to convert num to STRING

	str := strconv.Itoa(num)

	//then, we can now return ourstr,of type string
	fmt.Println("original Integer Number: ", num)
	fmt.Println("Already-converted String: ", str)

}

// FINDING Odd and Even Numbers
func evenOrOdd(num int) {

	//first range through number 1 to the num provided
	//...while identify both odd and even numbers in the range

	evenNumbers := 0
	oddNumbers := 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			evenNumbers++
			fmt.Printf("\n %v is EVEN\n", i)

		} else {
			oddNumbers++
			fmt.Printf("\n %v is ODD\n", i)
		}
	}

	//printing number of odd numbers in the lsit
	fmt.Printf("\n Total ODD numbers Encountered: %v \n", oddNumbers)

	//printing number of odd numbers in the lsit
	fmt.Printf("\nTotal EVEN numbers Encountered: %v\n", evenNumbers)

}

// defining a squareSum function
// Squares every number passed to it and returns the sum
func squareSum(num ...int) {
	sum := 0
	//receive the number(s) passed and range through them
	for i := 0; i < len(num); i++ {
		sqr := num[i] * num[i]
		sum += sqr
		fmt.Printf("\nSQUARE OF %v is %v\n", num[i], sqr)
	}
	fmt.Printf("\nSUM OF THEIR SQUARES: %v\n", sum)

}

// defining a strRepeat function
// accepting STRING and INT as parameter; then prints string INT number of times
func strRepeat(printStr string, intNoTimes int) {
	//ensure we print our string intNumber Of Time, intNoTimes
	for i := 0; i < intNoTimes; i++ {
		fmt.Println(printStr)
	}
}

// defining Grasshopper Summation
// returns sum of every number from 1 to num
func hopperSummation(number int) {
	sum := 0
	//first range through the numbers between 0 to the number passed to his function

	for i := 1; i <= number; i++ {
		sum += i
	}

	fmt.Printf("\nSum of integers from 1 to %v is %v\n", number, sum)
}

// Summing Number’s digits
func sumNumDigits(num int) {
	//first convert our num to a string
	str := strconv.Itoa(num)

	//then convert the string to a slice of runes
	runes := []rune(str)

	//now range through the string and sum up its data elements
	sum := 0
	for i := 0; i < len(runes); i++ {
		//convert every data element of the slice runes
		//first from a rune to a string
		strElement := string(runes[i])

		//then from string back to integer number
		intElement, _ := strconv.Atoi(strElement)

		//then finnaly sum up our converted int element
		//sum += int(runes[i])
		sum += intElement
	}
	fmt.Printf("\n sum of %v's digits: %v\n", num, sum)
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
	//////////////////
	/////////////////

	concert := "Reggae"
	str2Int(concert)
	////////////////////
	int2Str(55)
	evenOrOdd(31)
	strRepeat("love", 10)

	squareSum(2, 3, 4, 5, 10)
	hopperSummation(5)
	sumNumDigits(555)

}
