package castservice

import (
	"fmt"
	"movieapp/api/param"
	"movieapp/entity"
)

type Repository interface {
	GetAllCasts() ([]entity.Cast, error)
	GetCastById(int) (*entity.Cast, error)
	Create(entity.Cast) (*entity.Cast, error)
	Delete(int) (*entity.Cast, error)
	IsExist(int) (*bool, error)
}

type Service struct {
	castRepo Repository
}

func (s *Service) New(castRepo Repository) {
	s.castRepo = castRepo
}

func (s *Service) ListAllCasts() ([]entity.Cast, error) {
	all_cast, err := s.castRepo.GetAllCasts()
	if err != nil {
		return nil, fmt.Errorf("get all cast : %w\n", err)
	}
	return all_cast, nil
}

func (s *Service) GetCastById(Id int) (param.GetCastByIdResponse, error) {
	cast, err := s.castRepo.GetCastById(Id)
	if err != nil {
		return param.GetCastByIdResponse{}, fmt.Errorf("get cast by id : %w\n", err)
	} else if cast == nil {
		return param.GetCastByIdResponse{}, &ErrCastNotFound
	}
	return param.GetCastByIdResponse{Cast: cast}, nil
}

func (s *Service) Create(ccr param.CreateCastRequest) (*entity.Cast, error) {
	new_cast := entity.Cast{
		Name:   ccr.Name,
		Id:     -1,
		Movies: []int{},
	}
	cast, err := s.castRepo.Create(new_cast)
	if err != nil {
		return nil, fmt.Errorf("create cast : %w\n", err)
	}
	return cast, nil
}

func (s *Service) Delete(id int) (param.DeleteCastResponse, error) {
	ok, err := s.castRepo.IsExist(id)
	if err != nil {
		return param.DeleteCastResponse{}, fmt.Errorf("is cast exist : %w\n", err)
	}
	if !*ok {
		return param.DeleteCastResponse{}, &ErrCastNotFound
	}
	cast, err := s.castRepo.Delete(id)
	if err != nil {
		return param.DeleteCastResponse{}, fmt.Errorf("delete cast : %w\n", err)
	}
	return param.DeleteCastResponse{Cast: cast}, nil
}

func (s *Service) IsExist(id int) (*bool, error) {
	ok, err := s.castRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is cast exist : %w\n", err)
	}
	return ok, nil
}
