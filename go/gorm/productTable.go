package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open("sqlite3", "testdb.sqlite")
	if err != nil {
		panic("Not able to open db.")
	}
	defer db.Close()
	db.AutoMigrate(&Product{})
	db.Create(&Product{Code: "12", Price: 100})
	var pdt Product
	// db.LogMode(true)
	db.First(&pdt, "Code=?", "12")
	fmt.Println("pdt search: ", pdt.ID, pdt.Code, pdt.Price)

	db.Model(&pdt).Update("Price", 250)
	fmt.Println("pdt search: ", pdt.Code, pdt.Price)

	var pdts []Product
	db.Create(&Product{Code: "14", Price: 500})
	db.Find(&pdts)
	// fmt.Println(pdts)
	for index, pdtValue := range pdts {
		fmt.Println("Product: ", index, pdtValue.Code, pdtValue.Price)
		fmt.Println("------------")
	}

	db.Delete(&pdt)
	db.Delete(&pdts)
}
