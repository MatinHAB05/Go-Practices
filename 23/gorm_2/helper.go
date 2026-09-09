package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
)

func init() {
	Endcoder.SetIndent("", " ")

}
func Fatal(err error) {
	if err != nil {
		log.Println(0)
		log.Fatal(err)
	}
}
func FatalDB(db *gorm.DB) {
	if err := db.Error; err != nil {
		log.Println(0)
		log.Fatal(err)
	}
}

var jsonF, _ = os.Create("output.json")
var Endcoder = json.NewEncoder(jsonF)

func Show[T any](o T) {

	fmt.Println(Endcoder.Encode(o))
}
func ShowList[T any](o []T) {
	fmt.Println(Endcoder.Encode(o))
}

func Mig(db *gorm.DB, drop bool) {
	if drop {
		err := db.Migrator().DropTable(
			&CustomerLoan{},
			&CustomerAccount{},
			&Loan{},
			&Account{},
			&Branch{},
			&Bank{},
			&Customer{},
		)
		Fatal(err)

	}

	err := db.AutoMigrate(&Bank{})
	Fatal(err)

	err = db.AutoMigrate(&Branch{})
	Fatal(err)

	err = db.AutoMigrate(&Account{})
	Fatal(err)

	err = db.AutoMigrate(&Loan{})
	Fatal(err)

	err = db.AutoMigrate(&Customer{})
	Fatal(err)

	err = db.AutoMigrate(&CustomerAccount{})
	Fatal(err)

	err = db.AutoMigrate(&CustomerLoan{})
	Fatal(err)
}

func (b *Bank) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate")
	return nil
}
func (b *Bank) BeforeSave(tx *gorm.DB) error {
	log.Println("BeforeSave")
	return nil
}
func (b *Bank) BeforeUpdate(tx *gorm.DB) error {
	log.Println("BeforeUpdate")
	return nil
}
func (b *Bank) BeforeDelete(tx *gorm.DB) error {
	log.Println("BeforeDelete")
	return nil
}
func (b *Bank) AfterCreate(tx *gorm.DB) error {
	log.Println("AfterCreate")
	return nil
}
func (b *Bank) AfterSave(tx *gorm.DB) error {
	log.Println("AfterSave")
	return nil
}
func (b *Bank) AfterUpdate(tx *gorm.DB) error {
	log.Println("AfterUpdate")
	return nil
}
func (b *Bank) AfterDelete(tx *gorm.DB) error {
	log.Println("AfterDelete")
	return nil
}
func (b *Bank) BeforeFind(tx *gorm.DB) error {
	log.Println("BeforeUpdate")
	return nil
}
func (b *Bank) AfterFind(tx *gorm.DB) error {
	log.Println("AfterFind")
	return nil
}
