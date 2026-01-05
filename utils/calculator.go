package utils


type Calculator struct {
	LastResult int
}


func (c *Calculator) Multiply(a, b int) int {
	result := a * b
	c.LastResult = result
	return result
}
