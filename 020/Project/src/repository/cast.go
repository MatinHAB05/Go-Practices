package repository

import "movieapp/entity"

type CastRepository interface {
	GetAllCasts() ([]entity.Cast, error)
	GetCastById(int) (*entity.Cast, error)
	GetCastByIds([]int) ([]entity.Cast, error)
	Create(entity.Cast) (*entity.Cast, error)
	Delete(int) (*entity.Cast, error)
	IsExist(int) (*bool, error)
}
