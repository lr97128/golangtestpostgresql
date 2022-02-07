package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	URL      string = "10.188.88.68"
	PORT     int    = 5432
	USER     string = "root"
	PASSWORD string = "liurui97128224"
	DBNAME   string = "go"
)

type Product struct {
	ID    int     `sql:"product_no"`
	Name  string  `sql:"name"`
	Price float32 `sql:"price"`
}

func main() {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", URL, PORT, USER, PASSWORD, DBNAME)
	db, err := sql.Open("postgres", pgInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("连接postgresql成功")

	product, err := SelectProduct(db, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(product)

	// var product1 Product = Product{
	// 	ID:    5,
	// 	Name:  "HuaWei P50 pro",
	// 	Price: 7999,
	// }
	// err = InsertProduct(db, product1)
	// if err != nil {
	// 	panic(err)
	// }
	products, err := SelectProductsGreate(db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(products)
}

func InsertProduct(db *sql.DB, product Product) error {
	var sql string = "INSERT INTO Proto(product_no, name, price) VALUES($1, $2, $3)"
	result, err := db.Exec(sql, product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func SelectProduct(db *sql.DB, ProductNum int) (Product, error) {
	var product Product
	var sql string = "SELECT * from Proto WHERE product_no=$1"
	rows, err := db.Query(sql, ProductNum)
	if err != nil {
		return product, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&(product.ID), &(product.Name), &(product.Price))
		if err != nil {
			return product, err
		}
	}
	err = rows.Err()
	if err != nil {
		return product, err
	}
	return product, nil
}

func SelectProductsGreate(db *sql.DB, ProductNum int) ([]Product, error) {
	var products []Product
	var sql string = "SELECT * from Proto WHERE product_no>$1"
	rows, err := db.Query(sql, ProductNum)
	if err != nil {
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		var product Product
		err = rows.Scan(&(product.ID), &(product.Name), &(product.Price))
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, err
}
