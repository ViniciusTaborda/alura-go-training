package models

import (
	"server/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func GetAllProducts() []Product {

	dbConnection := db.GetDBConnection()

	defer dbConnection.Close()

	productsSelect, err := dbConnection.Query("SELECT * FROM PRODUCTS")

	if err != nil {
		panic(err)
	}

	product := Product{}
	products := []Product{}

	for productsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productsSelect.Scan(&name, &description, &price, &quantity, &id)

		if err != nil {
			panic(err)
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

		products = append(products, product)

	}

	return products
}

func InsertProduct(name, description string, price float64, quantity int) {

	dbConnection := db.GetDBConnection()
	defer dbConnection.Close()

	preparedStmt, err := dbConnection.Prepare(
		"INSERT INTO PRODUCTS(name, description, price, quantity) " +
			"VALUES($1, $2, $3, $4)",
	)

	if err != nil {
		panic(err)
	}

	preparedStmt.Exec(name, description, price, quantity)

	return
}

func DeleteProduct(productId string) {
	dbConnection := db.GetDBConnection()
	defer dbConnection.Close()

	preparedStmt, err := dbConnection.Prepare(
		"DELETE FROM PRODUCTS WHERE id=$1",
	)

	if err != nil {
		panic(err)
	}

	preparedStmt.Exec(productId)

	return
}

func EditProduct(productId string) Product {
	dbConnection := db.GetDBConnection()
	defer dbConnection.Close()

	productFromDb, err := dbConnection.Query(
		"SELECT * FROM PRODUCTS WHERE id=$1", productId,
	)

	if err != nil {
		panic(err)
	}

	product := Product{}

	for productFromDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productFromDb.Scan(&name, &description, &price, &quantity, &id)

		if err != nil {
			panic(err)
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Quantity = quantity

	}

	return product
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	dbConnection := db.GetDBConnection()
	defer dbConnection.Close()

	preparedStmt, err := dbConnection.Prepare(
		"UPDATE PRODUCTS SET " +
			"name=$1, description=$2, price=$3, quantity=$4 " +
			"WHERE id=$5",
	)

	if err != nil {
		panic(err)
	}

	preparedStmt.Exec(name, description, price, quantity, id)

	return
}
