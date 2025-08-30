package main

import "fmt"

type Persion struct {
	name string
	age  int
}

type Employee struct {
	Persion
	company string
}

func (e *Employee) setPrintInfo(name string, age int, company string) {
	e.name = name
	e.age = age
	e.company = company
}
func (e *Employee) PrintInfo() {
	fmt.Println("employee info: name->", e.name, "age->", e.age, "company->", e.company)
}

func main17() {
	employee := Employee{}
	employee.setPrintInfo("张三", 30, "alibaba")
	employee.PrintInfo()

}
