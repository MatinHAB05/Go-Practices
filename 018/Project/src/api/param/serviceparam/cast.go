package serviceparam

import "movieapp/entity"

type GetCastByIdInput struct {
	ID int `uri:"id"`
}

type RemoveCastByIdInput struct {
	ID int `json:"id"`
}

type CreateCastInput struct {
	Name string `json:"name"`
}

type DeleteCastResponse struct {
	Cast *entity.Cast `json:"casts"`
}
