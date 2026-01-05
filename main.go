package main

import (
	"fmt"

	"github.com/test-repo/test-repo-golang-support-v1/utils"
)

var (
	globalCounter int
	appName       string = "GraphTester"
)

func main() {
	result := utils.Add(10, 5)
	fmt.Printf("Initial Result: %d\n", result)

	subResult := utils.Subtract(10, 5)
	fmt.Printf("Subtracted Result: %d\n", subResult)

	calc := &utils.Calculator{}
	mulResult := calc.Multiply(10, 5)
	fmt.Printf("Multiplied Result: %d\n", mulResult)

	var op utils.Operator = calc
	interfaceResult := op.Multiply(6, 7)
	fmt.Printf("Interface Result: %d\n", interfaceResult)
	var num utils.Number = 100
	fmt.Printf("Number type: %d\n", num)

	fmt.Printf("Max Iterations: %d\n", utils.MaxIterations)
	fmt.Printf("Pi: %.5f\n", utils.Pi)

	globalCounter = 42
	fmt.Printf("App: %s, Counter: %d\n", appName, globalCounter)
}


