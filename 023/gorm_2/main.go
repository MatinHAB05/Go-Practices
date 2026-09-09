package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("testapp:testpass@tcp(:3308)/bankdb?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	Mig(db, false)
	db.Set("password", "123123")
	var banks []Bank
	db.Debug().Model(&Bank{}).Preload("Branches").Preload("Branches.Accounts").Preload("Branches.Accounts.CustomerAccounts").Find(&banks)
	Show(banks)
	// Show(		banks[0].Branches[0].Accounts[0].CustomerAccounts)
}
