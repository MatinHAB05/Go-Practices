package castservice

import (
	"fmt"
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
	"movieapp/repository"
	"movieapp/service/errorservice"
)

type Service struct {
	castRepo repository.CastRepository
}

func (s *Service) New(castRepo repository.CastRepository) {
	s.castRepo = castRepo
}

func (s *Service) ListAllCasts() ([]entity.Cast, error) {
	all_cast, err := s.castRepo.GetAllCasts()
	if err != nil {
		return nil, fmt.Errorf("get all cast : %w\n", err)
	}
	return all_cast, nil
}

func (s *Service) GetCastById(Id int) (*entity.Cast, error) {
	cast, err := s.castRepo.GetCastById(Id)
	if err != nil {
		return nil, fmt.Errorf("get cast by id-%d : %w\n", Id, err)
	} else if cast == nil {
		return nil, &errorservice.ErrCastNotFound
	}
	return cast, nil
}

func (s *Service) Create(ccr serviceparam.CreateCastInput) (*entity.Cast, error) {
	new_cast := entity.Cast{
		Name: ccr.Name,
		Id:   -1,
	}
	cast, err := s.castRepo.Create(new_cast)
	if err != nil {
		return nil, fmt.Errorf("create cast : %w\n", err)
	}
	return cast, nil
}

func (s *Service) Delete(id int) (*serviceparam.DeleteCastResponse, error) {
	ok, err := s.castRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is exist cast id-%d : %w\n", id, err)
	} else if !*ok {
		return nil, &errorservice.ErrCastNotFound
	}

	cast, err := s.castRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("delete cast id-%d : %w\n", id, err)
	}
	return &serviceparam.DeleteCastResponse{Cast: cast}, nil
}

func (s *Service) IsExist(id int) (*bool, error) {
	ok, err := s.castRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is cast id-%d exist : %w\n", id, err)
	}
	return ok, nil
}
