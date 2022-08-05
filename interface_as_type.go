package main

import "fmt"

type Manager struct {
	Name string
}

func (mgr Manager) UseProperty() string {
	return mgr.Name
}
func (mgr Manager) WithoutUseProperty() {
	fmt.Println("WithoutUseProperty")
}

type Employee interface {
	WithoutUseProperty()
	UseProperty() string
}

func main() {
	var newManager Manager
	var employeeInterface Employee

	employeeInterface = newManager
	name := employeeInterface.UseProperty()
	fmt.Println("UseProperty: ", name)
	employeeInterface.WithoutUseProperty()
	fmt.Println(employeeInterface == nil) // false

	newManager = Manager{"aaa"}
	employeeInterface = newManager
	name = employeeInterface.UseProperty()
	fmt.Println("UseProperty: ", name)
}
