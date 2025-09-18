package main

import "log"

//假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
//要求 ：
//定义一个 Book 结构体，包含与 books 表对应的字段。
//编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。

type Book struct {
	ID     uint    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func SelectBooksByPrice(db *sqlx.DB, price float64) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", price)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func main() {
	db, err := sqlx.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	books, err := SelectBooksByPrice(db, 50)
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		log.Printf("Book: %+v\n", book)
	}
}
