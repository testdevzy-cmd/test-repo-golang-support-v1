package main

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

func main() {
	result := utils.Add(10, 5)
	fmt.Printf("Initial Result: %d\n", result)
	
	diff := utils.Subtract(10, 5)
	fmt.Printf("Difference: %d\n", diff)

	quotient := utils.Divide(10, 5)
	fmt.Printf("Quotient: %d\n", quotient)

	rem := utils.Modulo(10, 3)
	fmt.Printf("Remainder: %d\n", rem)
}
