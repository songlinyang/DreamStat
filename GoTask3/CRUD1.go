package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Age       uint
	Grade     string
	CreatedAt time.Time      // Automatically managed by GORM for creation time
	UpdatedAt time.Time      // Automatically managed by GORM for update time
	DeleteAt  gorm.DeletedAt `gorm:"index"`
}
type StudentResult struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Age   uint
	Grade string
}

func main() {
	//1. 链接数据库
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//2.创建	Student表
	if db.Migrator().HasTable(&Student{}) == false {
		err := db.Migrator().CreateTable(&Student{}).Error()
		if err == "" {
			panic(err)
		}
	}
	//3.CUID操作
	//3.1编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	//3.2编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	var stu = Student{Age: 18}
	var stu_result = StudentResult{}
	db.Where("age > ?", stu.Age).Find(&stu).Scan(&stu_result)
	fmt.Println("查找结果：", stu_result)

	//3.3编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	var stu_result2 = StudentResult{}
	db.Model(&Student{}).Where("name=?", "张三").Update("grade", "四年级").Scan(&stu_result2)
	fmt.Println("更新结果：", stu_result2)

	//3.4编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。 //软删除
	var stu_result3 = StudentResult{}
	db.Where("age < ?", "15").Delete(&Student{}).Scan(&stu_result3)
	fmt.Println("删除后结果：", stu_result3)

	// 硬删除
	db.Unscoped().Where("age < ?", "15").Delete(&Student{}).Scan(&stu_result3)
	// DELETE FROM orders WHERE id=10;

}
