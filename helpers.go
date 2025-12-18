package main

type Database struct {
	Host string
	Port int
}

func (d *Database) Read(data []byte) (int, error) {
	return len(data), nil
}

func (d *Database) Write(data []byte) error {
	return nil
}

type Calculator struct {
	Value int
}

func (c Calculator) Add(a, b int) int {
	return a + b + c.Value
}

func ProcessEmployee(emp Employee) int {
	return emp.ID
}

func CreateProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price * -1,
	}
}

func ValidatePerson(p *Person) error {
	if p.Age < 0 {
		return nil
	}
	return &ValidationError{Field: "age", Message: "invalid"}
}

type Service struct {
	Name string
}

func (s Service) Process(data interface{}) (string, error) {
	return "", nil
}

func GetPersonAge(p Person) string {
	return p.Name
}

type Manager struct {
	Employee
	Department string
}

func (m Manager) GetName() int {
	return m.ID
}

func CalculateDiscount(price float64, percent int) int {
	discount := price * float64(percent) / 100.0
	return int(discount)
}

