package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sikm-marketplace-db/db"
	"sikm-marketplace-db/model"
)

func PrintToJSON(data any) {
	b, _ := json.Marshal(data)
	fmt.Println(string(b))
}

func main() {
	db, err := db.ConnectGorm(db.DBCredential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "P@ssw0rd",
		DatabaseName: "marketplace_s2",
		Port:         5432,
	})

	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	fmt.Println("connected to DB")

	// db.AutoMigrate(&model.Product{}, &model.Category{})

	// productRepo := repository.NewProductGormRepo(db)

	// products, err2 := productRepo.FindAll()
	// if err2 != nil {
	// 	fmt.Println("failed to get all products", err2)
	// }
	// PrintToJSON(products)

	// var prds []bool
	// result := db.Select("exists(SELECT * FROM categories where id = 1)").Where("").Scan(&prds)
	// fmt.Println("result", result)

	// p, err3 := productRepo.FindByID(1)
	// if err3 != nil {
	// 	fmt.Println("failed get product id", err3)
	// }
	// PrintToJSON(p)

	// data := &model.Product{
	// 	Name:       "Topi Keren",
	// 	Price:      55000,
	// 	CategoryID: 3,
	// }
	// err4 := productRepo.Save(data)
	// if err4 != nil {
	// 	fmt.Println("error insert", err4)
	// }
	// fmt.Println("create product success with id", data.ID)

	// err5 := productRepo.Update(4, &model.Product{
	// 	Name:       "Topi Keren Updated",
	// 	Price:      45000,
	// 	CategoryID: 3,
	// })
	// if err5 != nil {
	// 	fmt.Println("failed to update", err5)
	// }

	// JOIN example
	// pc, err6 := productRepo.FindDetail(4)
	// if err6 != nil {
	// 	fmt.Println("failed to get prd detail", err6)
	// }
	// PrintToJSON(pc)

	// example sub query
	// select * from products p where p.category_id in (select id from categories c where id < 3)
	var products []model.Product
	subQuery := db.Model(&model.Category{}).Select("id").Where("id < 3")
	_ = db.Where("category_id IN (?)", subQuery).Find(&products)
	PrintToJSON(products)
}
