package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	gorm.Model
	Name string
}

type Profile struct {
	gorm.Model
	UserID uint
	Name   string
	User   User
}

func main() {
	db, err := gorm.Open("sqlite3", "belongsTo.sqlite")
	if err != nil {
		panic("Can't open DB")
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Profile{})

	admin1 := User{Name: "admin"}
	db.Create(&admin1)
	adminProfile := Profile{User: admin1, Name: "admin profile"}
	db.Create(&adminProfile)
	var adminProfileSelect Profile
	db.Model(&admin1).Related(&adminProfileSelect)
	println(adminProfileSelect.User.Name, adminProfileSelect.Name)
	db.Delete(adminProfile)
	db.Delete(admin1)
}
