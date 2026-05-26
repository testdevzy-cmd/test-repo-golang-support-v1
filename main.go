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

	squared := utils.Square(8)
	fmt.Printf("Square of 8: %d\n", squared)

	runReviewSummaryDemo()
	runMathExtensionsDemo(calc)
}

func runReviewSummaryDemo() {
	table := utils.NewReviewSummaryTable(utils.ReviewSummaryVersion)
	table.AddRow(utils.ReviewSummaryRow{
		BaseStruct: utils.BaseStruct{ID: "1", CreatedAt: utils.TestVarTime},
		FilePath:   "utils/math.go",
		Line:       28,
		Severity:   utils.SeverityHigh,
		Message:    "Potential index out of bounds in loop condition",
		Category:   "bug",
		Confidence: 0.92,
	})
	table.AddRow(utils.ReviewSummaryRow{
		BaseStruct: utils.BaseStruct{ID: "2"},
		FilePath:   "utils/calculator.go",
		Line:       33,
		Severity:   utils.SeverityMedium,
		Message:    "Division without zero check",
		Category:   "reliability",
		Confidence: 0.81,
	})
	table.AddRow(utils.ReviewSummaryRow{
		FilePath:   "main.go",
		Line:       14,
		Severity:   utils.SeverityLow,
		Message:    "Consider extracting demo helpers",
		Category:   "maintainability",
		Confidence: 0.55,
	})

	fmt.Println(table.RenderMarkdown())
	counts := table.CountBySeverity()
	fmt.Printf("Severity counts: %+v\n", counts)
	fmt.Printf("High priority rows: %d\n", len(table.HighPriorityRows()))
}

func runMathExtensionsDemo(calc *utils.Calculator) {
	values := []int{3, 7, 11, 15}
	fmt.Println(utils.FormatStats("batch", values))

	div := utils.Divide(20, 4)
	fmt.Printf("Divide 20/4: %d\n", div)

	batch := utils.NewBatchCalculator("demo")
	results := batch.RunBatch([]string{"mul", "add", "div"}, 12, 3)
	fmt.Printf("Batch results: %v\n", results)

	fmt.Printf("Factorial(6): %d\n", utils.Factorial(6))
	fmt.Printf("IsPrime(17): %v\n", utils.IsPrime(17))

	weights := utils.DefaultSeverityWeights()
	fmt.Printf("Severity weights: %+v\n", weights)
	fmt.Printf("Calc history entries: %d\n", len(calc.History()))
}
