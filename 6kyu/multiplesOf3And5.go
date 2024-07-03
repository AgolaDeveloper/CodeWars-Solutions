package main

import "fmt"

func Multiple3And5(number int) int {
	//declaring variable sum that'll hold the total
	sum := 0

	for i := 1; i < number; i++ {
		//only incrementing our sum if the number is either a multiple of 3 or 5

		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	number := 50

	x := Multiple3And5(number)

	fmt.Println(x)
}
