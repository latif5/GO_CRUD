package models

import (
	"log"

	"github.com/Golang-CRUD/config"
	"github.com/Golang-CRUD/entities"
)

type ModelProduct struct {
}

func (*ModelProduct) FindAll() ([]entities.Product, error) {
	db, err := config.ConnDB()
	if err != nil {
		return nil, err
	} else {

		rows, err2 := db.Query("select * from product")
		if err2 != nil {
			return nil, err2
		} else {
			products := []entities.Product{}
			for rows.Next() {
				var product entities.Product
				rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil

		}

	}
}

func (*ModelProduct) FindById(id string) ([]entities.Product, error) {
	db, err := config.ConnDB()
	if err != nil {
		return nil, err
	} else {
		data, errSelect := db.Query("select * from product where id = ?", id)
		if errSelect != nil {
			return nil, errSelect
		} else {
			products := []entities.Product{}
			for data.Next() {
				var product entities.Product
				data.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Description)
				products = append(products, product)
			}
			return products, nil
		}
	}
}

func (*ModelProduct) AddData(product *entities.Product) bool {
	db, err := config.ConnDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into product(name, price, quantity, description) values(?, ?, ?,?)", product.Name, product.Price, product.Quantity, product.Description)
		if err2 != nil {
			log.Println(err2)
			return false
		} else {
			rowsAffected, errInsert := result.RowsAffected()
			if errInsert != nil {
				log.Println(errInsert)
				return false
			}
			return rowsAffected > 0
		}
	}
}

func (*ModelProduct) DeleteData(id int64) bool {
	db, err := config.ConnDB()
	log.Println(id)
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from product where id = ?", id)
		if err2 != nil {
			log.Println(err2)
			return false
		} else {
			rowsAffected, errDelete := result.RowsAffected()
			if errDelete != nil {
				log.Println(errDelete)
				return false
			}
			return rowsAffected > 0
		}
	}
}

func (*ModelProduct) UpdateData(product *entities.Product, id int64) bool {
	db, err := config.ConnDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("UPDATE product set name = ?, price = ?, quantity = ?, description = ? where id = ?", product.Name, product.Price, product.Quantity, product.Description, id)
		if err2 != nil {
			log.Println(err2)
			return false
		} else {
			rowsAffected, errInsert := result.RowsAffected()
			if errInsert != nil {
				log.Println(errInsert)
				return false
			}
			return rowsAffected > 0
		}
	}
}
