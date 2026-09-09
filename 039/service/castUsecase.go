package service

import (
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
)

type Cast interface {
	ListAllCasts() ([]entity.Cast, error)
	GetCastById(int) (*entity.Cast, error)

	Create(serviceparam.CreateCastInput) (*entity.Cast, error)

	Delete(int) (*serviceparam.DeleteCastResponse, error)

	IsExist(int) (*bool, error)
}
