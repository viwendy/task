package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

type Employee struct {
	Person
	EmployeeID uint
}

func (e *Employee) PrintInfo() *Employee {
	return e
}
func main() {
	em := &Employee{
		Person: Person{
			Name: "ABC",
			Age:  30,
		},
		EmployeeID: 12345,
	}
	info := em.PrintInfo()
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", info.Name, info.Age, info.EmployeeID)
}
