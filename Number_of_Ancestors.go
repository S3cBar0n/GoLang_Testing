package main

import "fmt"

func main() {
	var numAncestors int64 = 2
	var curYear int64 = 2021
	var userYear int64

	fmt.Print("Enter the year you want to calculate to: ")
	fmt.Scanf("%d", &userYear)

	for curYear >= userYear {
		fmt.Println(curYear, ":", numAncestors, "ancestors")

		numAncestors *= 2
		curYear -= 20
	}

}