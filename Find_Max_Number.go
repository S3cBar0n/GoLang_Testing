package main

import "fmt"

func main() {
	var maxSoFar int64
	var numValues int64
	var i int64 = 0

	fmt.Printf("Enter the total values you have: ")
	fmt.Scanf("%d\n", &numValues)

	for i < numValues {
		var currValue int64

		fmt.Printf("Input the next value: ")
		fmt.Scanf("%d\n", &currValue)

		if i == 0 {
			maxSoFar = currValue
			i += 1
		} else if currValue >= maxSoFar {
			maxSoFar = currValue
			i += 1
		} else {
			i += 1
		}
		fmt.Println("Max:", maxSoFar)
	}
}