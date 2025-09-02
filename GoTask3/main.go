package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 多对一
//type User struct {
//	gorm.Model
//	Name         string
//	CompanyRefer int
//	Company      Company `gorm:"foreignKey:CompanyRefer"` // 使用 CompanyRefer 作为外键
//}
//
//type Company struct {
//	ID   int
//	Name string
//}

// 多对一
//type User struct {
//	gorm.Model
//	Name      string
//	CompanyID string
//	Company   Company `gorm:"references:Code"` // 使用 Code 作为引用
//}
//
//type Company struct {
//	ID   int
//	Code string
//	Name string
//}

// 一对一
//type User struct {
//	gorm.Model
//	Name       string     `gorm:"index"`
//	CreditCard CreditCard `gorm:"foreignKey:UserName;references:Name"`
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number   string
//	UserName string
//}

//type User struct {
//	gorm.Model
//	Name      string
//	CompanyID string
//	Company   Company `gorm:"references:CompanyID"` // 使用 Company.CompanyID 作为引用
//}
//
//type Company struct {
//	CompanyID int
//	Code      string
//	Name      string
//}

// User 有多张 CreditCard，UserID 是外键
//type User struct {
//	gorm.Model
//	CreditCards []CreditCard
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number string
//	UserID uint
//}

//type User struct {
//	gorm.Model
//	CreditCard CreditCard `gorm:"foreignKey:UserName"` // 使用 UserName 作为外键
//}
//
//type CreditCard struct {
//	gorm.Model
//	Number   string
//	UserName string
//}

type User struct {
	gorm.Model
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func main2222() {
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Language{})

	db.Create(&User{
		Languages: []Language{{Name: "EN"}, {Name: "CH"}},
	})
}
