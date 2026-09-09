package mysqlmoviecast

import (
	"database/sql"
	"movieapp/api/param/repositoryparam"
	"movieapp/entity"
)

// GetMovieCastById(int) (*entity.MovieCast, error)

func (db *DB) Create(mc entity.MovieCast) (*entity.MovieCast, error) {
	conn := db.conn.Conn()

	_, err := conn.Exec(`
	INSERT INTO moviecast(movie_id,cast_id) VALUES
	(?,?)
	`, mc.MovieId, mc.CastId)

	if err != nil {
		return nil, err
	}

	return &mc, nil
}

func (db *DB) Delete(id int) (*entity.MovieCast, error) {
	conn := db.conn.Conn()

	mc, err := db.GetMovieCastById(id)
	if err != nil {
		return nil, err
	}
	if mc == nil {
		return nil, sql.ErrNoRows
	}

	_, err = conn.Exec("DELETE FROM moviecast WHERE id=?", id)
	if err != nil {
		return nil, err
	}
	return mc, nil

}

func (db *DB) DeleteByMovieIdCastId(movie_id, cast_id int) (*entity.MovieCast, error) {
	conn := db.conn.Conn()

	mc, err := db.GetMovieCastByMovieIdCastId(movie_id, cast_id)
	if err != nil {
		return nil, err
	}
	if mc == nil {
		return nil, sql.ErrNoRows
	}

	_, err = conn.Exec("DELETE FROM moviecast WHERE movie_id=? and cast_id=?", movie_id, cast_id)
	if err != nil {
		return nil, err
	}
	return mc, nil

}

func (db *DB) IsExist(id int) (*bool, error) {
	conn := db.conn.Conn()

	var exist bool
	err := conn.Get(&exist, `
		SELECT EXISTS(
			SELECT 1
			FROM moviecast
			WHERE id=?
		)
	`, id)

	if err != nil {
		return nil, err
	}

	return &exist, nil
}

func (db *DB) IsExistMovieCast(movie_id, cast_id int) (*bool, error) {
	conn := db.conn.Conn()

	var exist bool
	err := conn.Get(&exist, `
		SELECT EXISTS(
			SELECT 1
			FROM moviecast
			WHERE movie_id=? AND cast_id=?
		)
	`, movie_id, cast_id)

	if err != nil {
		return nil, err
	}

	return &exist, nil
}

func (db *DB) GetAllCastsOfMovie(movie_id int) ([]repositoryparam.GetAllCastsOfMovieResponse, error) {
	conn := db.conn.Conn()

	moviecasts := make([]repositoryparam.CastsOfMovieInput, 0)
	err := conn.Select(&moviecasts, `
	SELECT id,cast_id FROM moviecast
	WHERE moviecast.movie_id=?
	ORDER BY id ASC
	`, movie_id)

	if err != nil {
		return nil, err
	}

	final_mc := make([]repositoryparam.GetAllCastsOfMovieResponse, 0)
	for i := range moviecasts {
		final_mc = append(final_mc, repositoryparam.GetAllCastsOfMovieResponse{
			Id:     moviecasts[i].Id,
			CastId: moviecasts[i].CastId,
		})
	}

	return final_mc, nil

}

func (db *DB) GetAllMoviesOfCast(cast_id int) ([]repositoryparam.GetAllMoviesOfCastResponse, error) {
	conn := db.conn.Conn()

	moviecasts := make([]repositoryparam.MoviesOfCastInput, 0)
	err := conn.Select(&moviecasts, `
	SELECT id,movie_id FROM moviecast
	WHERE moviecast.cast_id=?
	ORDER BY id ASC
	`,
		cast_id)

	if err != nil {
		return nil, err
	}

	final_mc := make([]repositoryparam.GetAllMoviesOfCastResponse, 0)
	for i := range moviecasts {
		final_mc = append(final_mc, repositoryparam.GetAllMoviesOfCastResponse{
			Id:      moviecasts[i].Id,
			MovieId: moviecasts[i].MovieId,
		})
	}

	return final_mc, nil
}

func (db *DB) GetMovieCastById(id int) (*entity.MovieCast, error) {
	conn := db.conn.Conn()

	var mc repositoryparam.MovieCast
	err := conn.Get(&mc, `
	SELECT id,movie_id,cast_id FROM moviecast
	WHERE id=?`,
		id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &entity.MovieCast{
		Id:      mc.Id,
		MovieId: mc.MovieId,
		CastId:  mc.CastId,
	}, nil

}

func (db *DB) GetMovieCastByMovieIdCastId(movie_id, cast_id int) (*entity.MovieCast, error) {
	conn := db.conn.Conn()

	var mc repositoryparam.MovieCast
	err := conn.Get(&mc, `
	SELECT id,movie_id,cast_id FROM moviecast
	WHERE movie_id=? AND cast_id=?`,
		movie_id, cast_id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &entity.MovieCast{
		Id:      mc.Id,
		MovieId: mc.MovieId,
		CastId:  mc.CastId,
	}, nil
}

func (db *DB) GetAllMovieCast() ([]entity.MovieCast, error) {
	conn := db.conn.Conn()

	all := make([]repositoryparam.MovieCast, 0)
	err := conn.Select(&all, `
	SELECT id,movie_id,cast_id FROM moviecast
	ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}

	final_mc := make([]entity.MovieCast, 0)
	for i := range all {
		final_mc = append(final_mc, entity.MovieCast{
			Id:      all[i].Id,
			MovieId: all[i].MovieId,
			CastId:  all[i].CastId,
		})
	}

	return final_mc, nil
}
