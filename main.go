package main

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

func main() {
	result := utils.Add(10, 5)
	fmt.Printf("Initial Result: %d\n", result)

	subResult := utils.Subtract(10, 5)
	fmt.Printf("Subtracted Result: %d\n", subResult)

	calc := &utils.Calculator{}
	mulResult := calc.Multiply(10, 5)
	fmt.Printf("Multiplied Result: %d\n", mulResult)
}

