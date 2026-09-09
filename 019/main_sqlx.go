package main

import (
	"encoding/json"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	Id      int    `json:"id" db:"client_id"`
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Address string `json:"address" db:"address"`
	City    string `json:"city" db:"city"`
	State   string `json:"state" db:"state"`
}

func main() {
	db, err := sqlx.Connect("mysql", "testapp:testpass@tcp(:3308)/sql_invoicing")
	if err != nil {
		log.Fatalln(err)
	}
	// var clients []Client

	rows, err := db.PrepareNamed()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		// var mapp map[string]any = map[string]any{}
		var client Client
		err := rows.StructScan(client)

		// client.Id = int(mapp["client_id"].(int64))
		// client.Name = string(mapp["name"].([]uint8))
		// client.Address = string(mapp["address"].([]uint8))
		// client.City = string(mapp["city"].([]uint8))
		// client.State = string(mapp["state"].([]uint8))
		// client.Phone = string(mapp["phone"].([]uint8))

		Fatal(err)
		Final(client)
	}

}

func Final(final any) {
	x := json.NewEncoder(os.Stdout)
	x.SetIndent("", "  ")
	x.Encode(final)
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
	// log.Fatal()
}

// for rows.Next() {
// 	list, err := rows.SliceScan()
// 	var client Client

// 	client.Id = int(list[0].(int64))
// 	client.Name = string(list[1].([]uint8))
// 	client.Address = string(list[2].([]uint8))
// 	client.City = string(list[3].([]uint8))
// 	client.State = string(list[4].([]uint8))
// 	client.Phone = string(list[5].([]uint8))

// 	Fatal(err)
// 	Final(client)
// }

// 	// err = db.Get(&clients, "SELECT * FROM clients where client_id > ?", 1)
// // if err != nil {
// // 	log.Fatalln(err)
// // }

// 	var clients []Client
// err = db.Select(&clients, "SELECT * FROM clients where client_id > 5")
// if err != nil {
// 	log.Fatalln(err)
// }
// Final(clients)

// for rows.Next() {
// 	var mapp map[string]any = map[string]any{}
// 	err := rows.MapScan(mapp)
// 	var client Client

// 	client.Id = int(mapp["client_id"].(int64))
// 	client.Name = string(mapp["name"].([]uint8))
// 	client.Address = string(mapp["address"].([]uint8))
// 	client.City = string(mapp["city"].([]uint8))
// 	client.State = string(mapp["state"].([]uint8))
// 	client.Phone = string(mapp["phone"].([]uint8))

// 	Fatal(err)
// 	Final(client)
// }
