package models

import "github.com/vghessel/web_app/db"

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectWithDatabase()

	allProductsSelect, err := db.Query("SELECT * FROM products ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for allProductsSelect.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProductsSelect.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectWithDatabase()

	insertDataInDatabase, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertDataInDatabase.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectWithDatabase()

	deleteTheProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteTheProduct.Exec(id)
	defer db.Close()

}

func EditProduct(id string) Product {
	db := db.ConnectWithDatabase()

	databaseProduct, err := db.Query("SELECT * FROM products WHERE id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for databaseProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = databaseProduct.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity

	}

	defer db.Close()
	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectWithDatabase()

	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)
	defer db.Close()

}
