package moviecastservice

import (
	"errors"
	"fmt"
	"movieapp/api/param/serviceparam"
	"movieapp/entity"
	"movieapp/repository"
	"movieapp/service/errorservice"

	"github.com/go-sql-driver/mysql"
)

type Service struct {
	moviecastRepo repository.MovieCastRepository
}

func (s *Service) New(moviecastRepo repository.MovieCastRepository) {
	s.moviecastRepo = moviecastRepo
}

func (s *Service) GetAllMovieCasts() ([]entity.MovieCast, error) {
	all, err := s.moviecastRepo.GetAllMovieCast()
	if err != nil {
		return nil, fmt.Errorf("get all movie-casts : %w\n", err)
	}
	return all, nil
}

func (s *Service) GetAllCastsOfMovie(movie_id int) ([]serviceparam.GetAllCastsOfMovieRespnse, error) {
	casts, err := s.moviecastRepo.GetAllCastsOfMovie(movie_id)
	if err != nil {
		return nil, fmt.Errorf("get all casts of movie by movie_id id-%d : %w\n", movie_id, err)
	}
	final := make([]serviceparam.GetAllCastsOfMovieRespnse, 0)
	for i := range casts {
		final = append(final, serviceparam.GetAllCastsOfMovieRespnse{
			Id:     final[i].Id,
			CastId: final[i].CastId,
		})
	}
	return final, nil
}
func (s *Service) GetAllMoviesOfCast(cast_id int) ([]serviceparam.GetAllMoviesOfCastRespnse, error) {
	movies, err := s.moviecastRepo.GetAllMoviesOfCast(cast_id)
	if err != nil {
		return nil, fmt.Errorf("get all movies of cast by cast_id id-%d : %w\n", cast_id, err)
	}
	final := make([]serviceparam.GetAllMoviesOfCastRespnse, 0)
	for i := range movies {
		final = append(final, serviceparam.GetAllMoviesOfCastRespnse{
			Id:      final[i].Id,
			MovieId: final[i].MovieId,
		})
	}
	return final, nil
}

func (s *Service) GetMovieCastById(id int) (*entity.MovieCast, error) {
	mc, err := s.moviecastRepo.GetMovieCastById(id)
	if err != nil {
		return nil, fmt.Errorf("get movie-cast by id-%d : %w\n", id, err)
	} else if mc == nil {
		return nil, &errorservice.ErrMovieCastNotFound
	}
	return mc, nil
}

func (s *Service) GetMovieCastByMovieIdCastId(MovieId, CastId int) (*entity.MovieCast, error) {
	mc, err := s.moviecastRepo.GetMovieCastByMovieIdCastId(MovieId, CastId)
	if err != nil {
		return nil, fmt.Errorf("get movie-cast by movie-id id-%d & cast-id id-%d : %w\n", MovieId, CastId, err)
	} else if mc == nil {
		return nil, &errorservice.ErrMovieCastNotFound
	}
	return mc, nil
}

func (s *Service) Create(link serviceparam.LinkMovie2CastInput) (*entity.MovieCast, error) {
	ok, err := s.moviecastRepo.IsExistMovieCast(link.MovieId, link.CastId)
	if err != nil {
		return nil, fmt.Errorf("is movie-cast exist movie id-%d cast id-%d : %w\n", link.MovieId, link.CastId, err)
	} else if *ok {
		return nil, &errorservice.ErrMovieCastAlreadyLinked
	}
	moviecast := entity.MovieCast{
		Id:      -1,
		MovieId: link.MovieId,
		CastId:  link.CastId,
	}
	mc, err := s.moviecastRepo.Create(moviecast)
	if err != nil {
		var mysqlErr *mysql.MySQLError

		if errors.As(err, &mysqlErr) {
			if mysqlErr.Number == 1452 {
				return nil, &errorservice.ErrMovieCastForeignKeyFails
			}
		}
		return nil, fmt.Errorf("create movie-cast: %w\n", err)
	}
	return mc, nil
}

func (s *Service) DeleteByMovieIdCastId(MovieId int, CastId int) (*serviceparam.DeleteMovieCastResponse, error) {
	ok, err := s.moviecastRepo.IsExistMovieCast(MovieId, CastId)
	if err != nil {
		return nil, fmt.Errorf("is movie-cast exist movie id-%d cast id-%d : %w\n", MovieId, CastId, err)
	} else if !*ok {
		return nil, &errorservice.ErrMovieCastNotFound
	}

	mc, err := s.moviecastRepo.DeleteByMovieIdCastId(MovieId, CastId)
	if err != nil {
		return nil, fmt.Errorf("delete movie-cast by movie_id id-%d ,cast_id id-%d : %w\n", MovieId, CastId, err)
	}
	return &serviceparam.DeleteMovieCastResponse{MovieCast: mc}, nil
}

func (s *Service) Delete(id int) (*serviceparam.DeleteMovieCastResponse, error) {
	ok, err := s.moviecastRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is movie-cast exist id-%d : %w\n", id, err)
	} else if !*ok {
		return nil, &errorservice.ErrMovieCastNotFound
	}

	mc, err := s.moviecastRepo.Delete(id)
	if err != nil {
		return nil, fmt.Errorf("delete movie-cast id-%d : %w\n", id, err)
	}
	return &serviceparam.DeleteMovieCastResponse{MovieCast: mc}, nil
}

func (s *Service) IsExistMovieCast(MovieId, CastId int) (*bool, error) {
	ok, err := s.moviecastRepo.IsExistMovieCast(MovieId, CastId)
	if err != nil {
		return nil, fmt.Errorf("is movie-cast exist movie id-%d cast id-%d : %w\n", MovieId, CastId, err)
	}
	return ok, nil
}

func (s *Service) IsExist(id int) (*bool, error) {
	ok, err := s.moviecastRepo.IsExist(id)
	if err != nil {
		return nil, fmt.Errorf("is movie-cast exist id-%d : %w\n", id, err)
	}
	return ok, nil
}
