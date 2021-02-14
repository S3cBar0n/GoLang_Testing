package main

import "fmt"

func main() {
	var currC int8 = -10

	for currC <= 40 {
		var equivalentF float32 = float32((currC * 9.0 / 5.0) + 32.0)
		fmt.Println(currC, "is equivalent to", equivalentF, "F")
		currC += 5
	}
}