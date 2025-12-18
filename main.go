package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Constants
const (
	Pi      = 3.14159
	Version = "1.0.0"
)

const (
	StatusPending = iota
	StatusActive
	StatusCompleted
)

// Variables
var (
	globalCounter int
	mu            sync.Mutex
)

// Structs
type Person struct {
	Name string
	Age  int
}

// Embedded struct
type Employee struct {
	Person
	ID       int
	Position string
}

// Product struct for testing embeddings
type Product struct {
	ID          string
	Name        string
	Price       float64
	InStock     bool
	Tags        []string
	Metadata    map[string]interface{}
}

// Methods - Value receiver
func (p Person) GetName() string {
	return p.Name
}

// Methods - Pointer receiver
func (p *Person) SetAge(age int) {
	p.Age = age
}

// Interfaces
type Speaker interface {
	Speak() string
}

type Writer interface {
	Write([]byte) (int, error)
}

// Embedded interfaces
type ReadWriter interface {
	Reader
	Writer
}

type Reader interface {
	Read([]byte) (int, error)
}

// Custom types
type Status int
type ID string

// Error types
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Person implements Speaker interface
func (p Person) Speak() string {
	return fmt.Sprintf("Hello, I'm %s", p.Name)
}

// Functions
func Add(a, b int) int {
	return a + b
}

// Multiple return values
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Named return values
func Calculate(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}

// Variadic function
func Sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function as value
var Multiply = func(a, b int) int {
	return a * b
}

// Higher-order function
func ApplyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Closure
func CreateCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Generics
func Max[T comparable](a, b T) T {
	if a == b {
		return a
	}
	return a
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// Pointers
func Increment(n *int) {
	*n++
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

// Slices
func ProcessSlice(s []int) []int {
	result := make([]int, len(s))
	for i, v := range s {
		result[i] = v * 2
	}
	return result
}

// Maps
func ProcessMap(m map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		result[k] = v * 2
	}
	return result
}

// Type assertions
func AssertType(v interface{}) {
	if s, ok := v.(string); ok {
		fmt.Printf("String: %s\n", s)
	}
}

// Type switch
func TypeSwitch(v interface{}) string {
	switch v.(type) {
	case int:
		return "integer"
	case string:
		return "string"
	case bool:
		return "boolean"
	default:
		return "unknown"
	}
}

// Defer
func DeferExample() {
	defer fmt.Println("This runs last")
	fmt.Println("This runs first")
}

// Panic and Recover
func PanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v\n", r)
		}
	}()
	panic("something went wrong")
}

// Goroutines
func GoroutineExample() {
	ch := make(chan string, 1)
	go func() {
		ch <- "Hello from goroutine"
	}()
	msg := <-ch
	fmt.Println(msg)
}

// Channels
func ChannelExample() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for v := range ch {
		fmt.Printf("Received: %d\n", v)
	}
}

// Select
func SelectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	select {
	case msg := <-ch1:
		fmt.Printf("Select: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("Select: %s\n", msg)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Select timeout")
	}
}

// WaitGroup
func WaitGroupExample() {
	var wg sync.WaitGroup
	results := make([]int, 0)
	var mu sync.Mutex

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			results = append(results, id)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("WaitGroup results: %v\n", results)
}

// Mutex
func MutexExample() {
	var counter int
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Printf("Counter: %d\n", counter)
}

// Context
func ContextExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("Context timeout")
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Should not reach here")
	}
}

// Main function
func main() {
	fmt.Println("=== Go Language Constructs Demo ===")

	// Variables and constants
	fmt.Println("1. Constants and Variables:")
	fmt.Printf("Pi: %.5f, Version: %s\n", Pi, Version)
	fmt.Printf("StatusPending: %d, StatusActive: %d\n\n", StatusPending, StatusActive)

	// Structs
	fmt.Println("2. Structs:")
	p := Person{Name: "Alice", Age: 40}
	fmt.Printf("Person: %+v\n", p)
	fmt.Printf("Name: %s\n", p.GetName())
	p.SetAge(31)
	fmt.Printf("Age after SetAge: %d\n\n", p.Age)

	// Embedded structs
	fmt.Println("3. Embedded Structs:")
	emp := Employee{
		Person:   Person{Name: "Bob", Age: 25},
		ID:       123,
		Position: "Developer",
	}
	fmt.Printf("Employee: %+v\n", emp)
	fmt.Printf("Employee Name: %s\n\n", emp.GetName())

	// Interfaces
	fmt.Println("4. Interfaces:")
	var speaker Speaker = p
	fmt.Printf("Speaker: %s\n\n", speaker.Speak())

	// Functions
	fmt.Println("5. Functions:")
	fmt.Printf("Add(5, 3): %d\n", Add(5, 3))
	sum, prod := Calculate(4, 5)
	fmt.Printf("Calculate(4, 5): sum=%d, product=%d\n", sum, prod)
	fmt.Printf("Sum(1,2,3,4,5): %d\n", Sum(1, 2, 3, 4, 5))
	fmt.Printf("Multiply(3, 4): %d\n", Multiply(3, 4))
	fmt.Printf("ApplyOperation(10, 5, Add): %d\n\n", ApplyOperation(10, 5, Add))

	// Closures
	fmt.Println("6. Closures:")
	counter := CreateCounter()
	fmt.Printf("Counter: %d, %d, %d\n\n", counter(), counter(), counter())

	// Generics
	fmt.Println("7. Generics:")
	fmt.Printf("Max(10, 20): %d\n", Max(10, 20))
	stack := &Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	fmt.Printf("Stack Pop: %d\n\n", stack.Pop())

	// Pointers
	fmt.Println("8. Pointers:")
	x := 10
	Increment(&x)
	fmt.Printf("After Increment: %d\n", x)
	personPtr := NewPerson("Charlie", 35)
	fmt.Printf("Person pointer: %+v\n\n", *personPtr)

	// Slices
	fmt.Println("9. Slices:")
	slice := []int{1, 2, 3, 4, 5}
	processed := ProcessSlice(slice)
	fmt.Printf("Original: %v, Processed: %v\n\n", slice, processed)

	// Maps
	fmt.Println("10. Maps:")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	processedMap := ProcessMap(m)
	fmt.Printf("Original: %v, Processed: %v\n\n", m, processedMap)

	// Type operations
	fmt.Println("11. Type Operations:")
	AssertType("hello")
	fmt.Printf("TypeSwitch(42): %s\n", TypeSwitch(42))
	fmt.Printf("TypeSwitch(\"test\"): %s\n\n", TypeSwitch("test"))

	// Defer
	fmt.Println("12. Defer:")
	DeferExample()
	fmt.Println()

	// Panic and Recover
	fmt.Println("13. Panic and Recover:")
	PanicRecover()
	fmt.Println()

	// Goroutines and Channels
	fmt.Println("14. Goroutines and Channels:")
	GoroutineExample()
	ChannelExample()
	fmt.Println()

	// Select
	fmt.Println("15. Select:")
	SelectExample()
	fmt.Println()

	// WaitGroup
	fmt.Println("16. WaitGroup:")
	WaitGroupExample()

	// Mutex
	fmt.Println("17. Mutex:")
	MutexExample()
	fmt.Println()

	// Error handling
	fmt.Println("18. Error Handling:")
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Divide(10, 2): %.2f\n", result)
	}

	_, err = Divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	customErr := &ValidationError{Field: "email", Message: "invalid format"}
	fmt.Printf("Custom error: %v\n\n", customErr)

	fmt.Println("19. Utils Testing:")
	company := Company{Name: "TechCorp", Industry: "Software"}
	var companySpeaker Speaker = company
	fmt.Printf("Company speaker: %s\n", companySpeaker.Speak())

	person := Person{Name: "David", Age: 28}
	formatted := FormatPerson(person)
	fmt.Printf("Formatted: %s\n", formatted)
	fmt.Println()

	fmt.Println("20. Helpers Testing:")
	db := &Database{Host: "localhost", Port: 5432}
	var rw ReadWriter = db
	buffer := make([]byte, 10)
	rw.Read(buffer)
	rw.Write(buffer)

	calc := Calculator{Value: 10}
	calcResult := ApplyOperation(5, 3, calc.Add)
	fmt.Printf("Calculator result: %d\n", calcResult)

	empID := ProcessEmployee(emp)
	fmt.Printf("Employee ID: %d\n", empID)

	product := CreateProduct("Laptop", 999.99)
	fmt.Printf("Product: %+v\n", product)

	validErr := ValidatePerson(&p)
	if validErr != nil {
		fmt.Printf("Validation error: %v\n", validErr)
	}

	age := GetPersonAge(p)
	fmt.Printf("Person age: %d\n", age)

	mgr := Manager{
		Employee:   emp,
		Department: "Engineering",
	}
	name := mgr.GetName()
	fmt.Printf("Manager name: %s\n", name)
	fmt.Println()

	fmt.Println("21. Processors Testing:")
	handler := DataHandler{Name: "main"}
	processedData, procErr := handler.Process(nil, "test data")
	if procErr == nil {
		fmt.Printf("Processed: %s\n", processedData)
	}

	cache := &Cache{}
	cache.Set("key1", "value1", 60)
	val := cache.Get("key1")
	fmt.Printf("Cache value: %v\n", val)
	cache.Delete("key1")

	logger := Logger{prefix: "APP"}
	logger.Info("test message", "arg1")

	transformedData, transformErr := TransformData("input")
	if transformErr == nil {
		fmt.Printf("Transformed: %v\n", transformedData)
	}

	filtered := FilterData([]interface{}{1, 2, 3})
	fmt.Printf("Filtered: %v\n", filtered)

	trans := Transformer{config: "default"}
	transResult, transCode := trans.Transform("data")
	fmt.Printf("Transform result: %v (code: %d)\n", transResult, transCode)
	fmt.Println()

	fmt.Println("=== Demo Complete ===")
}

