package main

import "fmt"

func main() {
	var dashes int8  // 8bit numbers only -128 to 128

	fmt.Print("How many dashes: ")
	fmt.Scan(&dashes)

	for dashes > 0 {
		fmt.Print("-")
		dashes -= 1
	}

}