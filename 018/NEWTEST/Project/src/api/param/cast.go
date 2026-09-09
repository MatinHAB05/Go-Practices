package param

import "movieapp/entity"

type GetCastByIdRequest struct {
	ID int `uri:"id" binding:"required"`
}

type GetCastByIdResponse struct {
	Cast *entity.Cast
}

type CreateCastRequest struct {
	Name string `json:"name" binding:"omitempty,castName"`
}

type GetCastWithFilterRequest struct {
	Name string `form:"name" binding:"omitempty,castName"`
}

type DeleteCastResponse struct {
	Cast *entity.Cast
}
