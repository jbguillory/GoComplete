package main

import "fmt"

func main() {
	for _, num := range twoHands() {
		isEven := num%2 == 0
		fmt.Printf("%d is even: %t\n", num, isEven)
	}
}

func twoHands() []int {
	bothHands := make([]int, 10)
	bh := bothHands
	for i := range bh {
		bh[i] = i + 1
	}
	return bothHands
}
