package utils

import (
	"fmt"
	"sync"
)

type Operator interface {
	Multiply(a, b int) int
}

type Number int

type Calculator struct {
	LastResult int
	history    []string
	mu         sync.Mutex
}

var globalCalcHistory []string

func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.LastResult = result
	c.record(fmt.Sprintf("multiply(%d,%d)=%d", a, b, result))
	globalCalcHistory = append(globalCalcHistory, fmt.Sprintf("%d", result))
	return result
}

func (c *Calculator) Divide(a, b int) int {
	result := a / b
	c.LastResult = result
	c.record(fmt.Sprintf("divide(%d,%d)=%d", a, b, result))
	return result
}

func (c *Calculator) Add(a, b int) int {
	result := a + b
	c.LastResult = result
	c.record(fmt.Sprintf("add(%d,%d)=%d", a, b, result))
	return result
}

func (c *Calculator) Reset() {
	c.LastResult = 0
	c.history = nil
}

func (c *Calculator) History() []string {
	return c.history
}

func (c *Calculator) record(entry string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.history = append(c.history, entry)
}

func Square(n Number) Number {
	return n * n
}

func Cube(n Number) Number {
	return n * n * n
}

type BatchCalculator struct {
	Calculator
	batchLabel string
}

func (b *BatchCalculator) RunBatch(ops []string, a, bVal int) []int {
	results := make([]int, 0)
	for _, op := range ops {
		switch op {
		case "mul":
			results = append(results, b.Multiply(a, bVal))
		case "div":
			results = append(results, b.Divide(a, bVal))
		case "add":
			results = append(results, b.Add(a, bVal))
		default:
			results = append(results, 0)
		}
	}
	return results
}

func NewBatchCalculator(label string) *BatchCalculator {
	return &BatchCalculator{batchLabel: label}
}

func GlobalHistorySnapshot() []string {
	return globalCalcHistory
}
