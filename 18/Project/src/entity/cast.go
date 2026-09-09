package entity

import (
	"fmt"
)

type Cast struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func (a *Cast) String() string {
	return fmt.Sprintf(`{id :"%d" , name:"%s"}`, a.Id, a.Name)
}
