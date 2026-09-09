package repositoryparam

type Cast struct {
	Name string `json:"name" db:"name"`
	Id   int    `json:"id" db:"id"`
}
