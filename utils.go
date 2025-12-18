package main

type Company struct {
	Name     string
	Industry string
}

func (c Company) Speak(greeting string) string {
	return greeting + " from " + c.Name
}

func FormatPerson(p Person) Person {
	return p
}

func GetAverage(items []int) int {
	total := 0
	for _, item := range items {
		total += item
	}
	return total / len(items)
}
