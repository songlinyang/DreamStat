package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 注册 sqlite 驱动（必须）
)

type EmployeesResult struct {
	ID         uint
	Name       string
	Department string
	Salary     float64
}

// 题目1：使用SQL扩展库进行查询
func main6() {
	var (
		db  *sqlx.DB
		err error
		dsn = "identifier.sqlite"
	)
	// open and connect at the same time:
	db, err = sqlx.Connect("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	//插入数据
	//for i := 0; i < 10; i++ {
	//	employeesName := fmt.Sprint("技术员", i)
	//	department := "技术部"
	//	salary := float64(rand.Intn(100000))
	//	employees := map[string]interface{}{"name": employeesName, "department": department, "salary": salary}
	//	result, err := db.NamedExec(`INSERT INTO employees (name,department,salary) VALUES (:name, :department,:salary)`, employees)
	//	if err != nil {
	//		panic(err)
	//	}
	//	id, _ := result.LastInsertId()
	//	fmt.Println(id)
	//}
	//编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
	employ := EmployeesResult{Department: "技术部"}
	results := []EmployeesResult{}
	nstmt, err := db.PrepareNamed(`SELECT id,name,department,salary FROM employees WHERE department = :department`)
	err = nstmt.Select(&results, employ)
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	nstmt2, err2 := db.PrepareNamed(`SELECT id,name,department,salary FROM employees WHERE department = :department AND salary = (select max(salary) from employees where department=:department)`)
	err2 = nstmt2.Select(&results, employ)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(results)
}
