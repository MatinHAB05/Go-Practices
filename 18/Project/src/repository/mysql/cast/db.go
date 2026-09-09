package mysqlcast

import mysqldb "movieapp/repository/mysql"

type DB struct {
	conn *mysqldb.MySQLDB
}

func New(conn *mysqldb.MySQLDB) *DB {
	return &DB{
		conn: conn,
	}
}
