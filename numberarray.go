package main

import "fmt"

type numbers []int

func newNumberArray() numbers {
	numbersArray := numbers{1,2,3,4,5,6,7,8,9,10}
	newNumbers := numbers{}

	for _, number := range numbersArray {
		if number % 2 == 0 {
			fmt.Println(number, "is even")
			newNumbers = append(newNumbers, number )
		} else {
			fmt.Println(number, "is odd")
			newNumbers = append(newNumbers, number )
		}

	}
		return newNumbers
}
