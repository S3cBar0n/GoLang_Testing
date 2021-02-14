package main

import "fmt"

func main() {
	const interestRate float64 = 0.05
	var initialSavings float64
	var year int16 = 0

	fmt.Print("Enter your initial savings: ")
	fmt.Scanf("%f", &initialSavings)

	var currSavings float64 = initialSavings  // This is here because this needs to be declared after the scan to pull our users input
	for year < 10 {
		fmt.Println(currSavings)
		currSavings = currSavings + (currSavings * interestRate)
		year += 1
	}

}
