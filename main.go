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
}
