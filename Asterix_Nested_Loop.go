package main

import "fmt"

func main() {
	var numAsterix = []int8{1, 4, 7, 9, 12, -4}  // Array that contains my values

	for num := range numAsterix {
		if numAsterix[num] >= 0 {
			fmt.Print(numAsterix[num], ": ")
			var iterations int8 = 0
			for iterations < numAsterix[num] {
				fmt.Print("*")
				iterations += 1
			}
			fmt.Println()
		} else {
			break
		}

	}
}
