package main

import (
	"errors"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Accounts struct {
	gorm.Model
	Name    string
	Balance int
}
type Transactions struct {
	gorm.Model
	FromAccountID uint
	ToAccountID   uint
	Amount        int
}

func main() {
	//1. 链接数据库
	db, err := gorm.Open(sqlite.Open("identifier.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//2.创建	Student表
	db.AutoMigrate(&Accounts{})
	db.AutoMigrate(&Transactions{})
	//创建A，B账户
	accounts := []Accounts{{Name: "A", Balance: 100}, {Name: "B", Balance: 100}}
	err1 := db.Create(&accounts).Error
	if err1 != nil {
		panic(err1)
	}
	var a, b Accounts
	db.Debug().Where("name=?", "A").First(&a)
	db.Debug().Where("name=?", "B").First(&b)
	//3.事务
	resErr := db.Transaction(func(tx *gorm.DB) error {
		//转账100元前，先判断A是否足额
		if err := tx.Where("id = ?", a.ID).First(&a).Error; err != nil {
			return err
		}
		if a.Balance < 100 {

			return errors.New(a.Name + " balance is less than 100")
		}
		if err := tx.Where("id = ?", b.ID).First(&b).Error; err != nil {
			return err
		}
		// 额度足够，开始转账
		// A减去100
		if err := tx.Model(&Accounts{}).Where("name=?", "A").Update("balance", a.Balance-100).Error; err != nil {
			return err
		}
		// B加上100
		if err := tx.Debug().Model(&Accounts{}).Select("name", "balance").Where("name=?", "B").Update("balance", b.Balance+100).Error; err != nil {
			return err
		}
		// 记录转账交易
		if err := tx.Debug().Create(&Transactions{Amount: 100, FromAccountID: a.ID, ToAccountID: b.ID}).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务

		return nil
	})
	if resErr != nil {
		fmt.Println(resErr)
	} else {
		fmt.Println("转账成功!")
	}
}
