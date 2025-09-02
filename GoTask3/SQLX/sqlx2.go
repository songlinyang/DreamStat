package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3" // 注册 sqlite 驱动（必须）
)

type Books struct {
	ID     uint
	Title  string
	Author string
	Price  int
}

func main() {
	db, err := sqlx.Connect("sqlite3", "identifier.sqlite")
	if err != nil {
		panic(err)
	}
	//插入测试数据
	//for i := 0; i < 10; i++ {
	//	title := fmt.Sprint("书名", i)
	//	author := fmt.Sprint("侠名", i)
	//	price := rand.Intn(100)
	//	book := Books{Title: title, Author: author, Price: price}
	//	result, err := db.NamedExec(`INSERT INTO books (title,author,price) VALUES (:title, :author,:price)`, book)
	//	if err != nil {
	//		panic(err)
	//	}
	//	id, _ := result.LastInsertId()
	//	fmt.Println(id)
	//}
	var booksResult []Books
	p := Books{Price: 50}
	nstmt, err1 := db.PrepareNamed(`select id,title,author,price from books where price > :price`)
	if err1 != nil {
		panic(err1)
	}
	err2 := nstmt.Select(&booksResult, p)
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(booksResult)

}
