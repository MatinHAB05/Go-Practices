package bindparam

type GetCastByIdRequest struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCastRequest struct {
	Name string `json:"name" binding:"omitempty,castName"`
}

type RemoveCastByIdRequest struct {
	ID int `json:"id" binding:"required"`
}

type GetCastWithFilterRequest struct {
	Name string `form:"name" binding:"omitempty,castName"`
}

// type GetCastByIdResponse struct {
// 	Name   string         `json:"name"`
// 	Id     int            `json:"id"`
// 	Movies []entity.Movie `json:"movies"`
// }

//
// type DeleteCastResponse struct {
// 	Cast *entity.Cast
// }
