package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	conn, err := sql.Open("mysql", "testapp:testpass@tcp(:3308)/testdb")
// 	defer conn.Close()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	err = conn.Ping()
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	new_names := []string{"A", "B", "C", "D", "E", "F"}
// 	new_age := []int{1, 2, 3, 4, 5, 6, 7}

// 	stmt, err := conn.Prepare(`
// 	INSERT INTO User (name ,age) values
// 	('mamad',123),
// 	('saeed',123),
// 	('jasem',123),
// 	(?,?),
// 	(?,?),
// 	(?,?),
// 	(?,?),
// 	(?,?)

// 	`)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// stmt.

// 	// row := conn.QueryRow("SELECT * FROM User WHERE id=")
// 	// if err != nil {
// 	// 	log.Println(err)
// 	// 	return
// 	// }
// 	// var (
// 	// 	id   int
// 	// 	name string
// 	// 	age  int
// 	// )
// 	// fmt.Println(row.Err())
// 	// fmt.Println(row.Scan(&id, &name, &age))
// 	// fmt.Println(" === ", id, name, age)
// 	// // fmt.Println(rows)
// 	// // fmt.Println(rows.Columns())
// 	// // fmt.Println(rows.ColumnTypes())
// 	// // for i := 1; rows.Next(); i++ {
// 	// // 	var (
// 	// // 		id   int
// 	// // 		name string
// 	// // 		age  int
// 	// // 	)
// 	// // 	rows.Scan(&id, &name, &age)
// 	// // 	fmt.Println(i, " === ", id, name, age)
// 	// // 	rows.
// 	// // }
// 	fmt.Println("FIN")
// }


