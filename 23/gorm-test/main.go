package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID    int
	Name  string
	Age   int
	Phone string
}

func main() {

	dsn := "testapp:testpass@tcp(:3308)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	db.AutoMigrate(&User{})

	sql := "SELECT * FROM users WHERE name = ? AND age > ?"

	result := db.Dialector.Explain(sql, "Ali", 18)

	fmt.Println(result)

	// db.First()
	// db.Last()
	// db.Take()
	// db.Select()
	// db.Order()
	// db.UpdateColumn()پ
	// db.Preload()
	// db.Exec().Scan().Group().Get()
	db.Set().Omit	
}

