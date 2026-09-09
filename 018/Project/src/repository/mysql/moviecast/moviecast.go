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

	row := conn.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM moviecast
			WHERE id=?
		)
	`, id)

	var exist bool
	err := row.Scan(&exist)
	if err != nil {
		return nil, err
	}

	return &exist, nil
}

func (db *DB) IsExistMovieCast(movie_id, cast_id int) (*bool, error) {
	conn := db.conn.Conn()

	row := conn.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM moviecast
			WHERE movie_id=? AND cast_id=?
		)
	`, movie_id, cast_id)

	var exist bool
	err := row.Scan(&exist)
	if err != nil {
		return nil, err
	}

	return &exist, nil
}

func (db *DB) GetAllCastsOfMovie(movie_id int) ([]repositoryparam.GetAllCastsOfMovieRespnse, error) {
	conn := db.conn.Conn()
	rows, err := conn.Query(`
	SELECT id,cast_id FROM moviecast
	WHERE moviecast.movie_id=?
	ORDER BY id ASC
	`,
		movie_id)

	if err != nil {
		return nil, err
	}

	moviecasts := make([]repositoryparam.GetAllCastsOfMovieRespnse, 0)
	for rows.Next() {
		var mc repositoryparam.GetAllCastsOfMovieRespnse = repositoryparam.GetAllCastsOfMovieRespnse{}
		rows.Scan(&mc.Id, &mc.CastId)
		moviecasts = append(moviecasts, mc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return moviecasts, nil

}
func (db *DB) GetAllMoviesOfCast(cast_id int) ([]repositoryparam.GetAllMoviesOfCastRespnse, error) {
	conn := db.conn.Conn()
	rows, err := conn.Query(`
	SELECT id,movie_id FROM moviecast
	WHERE moviecast.cast_id=?
	ORDER BY id ASC
	`,
		cast_id)

	if err != nil {
		return nil, err
	}

	moviecasts := make([]repositoryparam.GetAllMoviesOfCastRespnse, 0)
	for rows.Next() {
		var mc repositoryparam.GetAllMoviesOfCastRespnse = repositoryparam.GetAllMoviesOfCastRespnse{}
		rows.Scan(&mc.Id, &mc.MovieId)
		moviecasts = append(moviecasts, mc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return moviecasts, nil

}

func (db *DB) GetMovieCastById(id int) (*entity.MovieCast, error) {
	conn := db.conn.Conn()
	row := conn.QueryRow(`
	SELECT id,movie_id,cast_id FROM moviecast
	WHERE id=?`,
		id)

	var mc entity.MovieCast
	err := row.Scan(&mc.Id, &mc.MovieId, &mc.CastId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &mc, nil

}

func (db *DB) GetMovieCastByMovieIdCastId(movie_id, cast_id int) (*entity.MovieCast, error) {
	conn := db.conn.Conn()
	row := conn.QueryRow(`
	SELECT id,movie_id,cast_id FROM moviecast
	WHERE movie_id=? AND cast_id=?`,
		movie_id, cast_id)

	var mc entity.MovieCast
	err := row.Scan(&mc.Id, &mc.MovieId, &mc.CastId)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}
	return &mc, nil
}

func (db *DB) GetAllMovieCast() ([]entity.MovieCast, error) {
	conn := db.conn.Conn()
	rows, err := conn.Query(`
	SELECT id,movie_id,cast_id FROM moviecast
	ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	all := make([]entity.MovieCast, 0)
	for rows.Next() {
		var mc entity.MovieCast
		rows.Scan(&mc.Id, &mc.MovieId, &mc.CastId)
		all = append(all, mc)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return all, nil
}
