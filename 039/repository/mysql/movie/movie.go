package mysqlmovie

import (
	"database/sql"
	"movieapp/api/param/repositoryparam"
	"movieapp/entity"

	"github.com/jmoiron/sqlx"
)

func (db *DB) GetAllMovies() ([]entity.Movie, error) {
	conn := db.conn.Conn()

	movies := make([]repositoryparam.Movie, 0)

	err := conn.Select(&movies, `
	SELECT id,title,quality,release_year FROM movie
	ORDER BY id ASC
	`)

	if err != nil {
		return nil, err
	}
	final_movies := make([]entity.Movie, 0)
	for i := range movies {
		final_movies = append(final_movies, entity.Movie{
			Title:       movies[i].Title,
			ReleaseYear: movies[i].ReleaseYear,
			Quality:     movies[i].Quality,
			Id:          movies[i].Id,
		})
	}
	return final_movies, nil
}

func (db *DB) GetMovieById(id int) (*entity.Movie, error) {
	conn := db.conn.Conn()

	var m repositoryparam.Movie
	err := conn.Get(&m, `
		SELECT id, title, quality, release_year
		FROM movie
		WHERE id=?
	`, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entity.Movie{
		Title:       m.Title,
		ReleaseYear: m.ReleaseYear,
		Quality:     m.Quality,
		Id:          m.Id,
	}, nil
}

func (db *DB) Create(m entity.Movie) (*entity.Movie, error) {
	conn := db.conn.Conn()

	res, err := conn.Exec(`
		INSERT INTO movie (title, quality, release_year)
		VALUES (?, ?, ?)
	`,
		m.Title,
		m.Quality,
		m.ReleaseYear,
	)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	m.Id = int(id)

	return &m, nil
}

func (db *DB) Delete(id int) (*entity.Movie, error) {
	conn := db.conn.Conn()

	movie, err := db.GetMovieById(id)
	if err != nil {
		return nil, err
	}

	if movie == nil {
		return nil, sql.ErrNoRows
	}

	_, err = conn.Exec("DELETE FROM movie WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (db *DB) IsExist(id int) (*bool, error) {
	conn := db.conn.Conn()

	var exists bool

	err := conn.Get(&exists, `
		SELECT EXISTS(
			SELECT id FROM movie WHERE id=?
		)
	`, id)

	if err != nil {
		return nil, err
	}

	return &exists, nil
}

func (db *DB) GetMovieByIds(ids []int) ([]entity.Movie, error) {
	conn := db.conn.Conn()

	q, args, err := sqlx.In(`
	SELECT id,title,release_year,quality
	FROM movie
	WHERE id IN (?)
	ORDER BY id ASC`, ids)
	if err != nil {
		return nil, err
	}

	movies := make([]repositoryparam.Movie, 0)
	err = conn.Select(&movies, q, args...)

	if err != nil {
		return nil, err
	}

	final_movies := make([]entity.Movie, 0)
	for i := range movies {
		final_movies = append(final_movies, entity.Movie{
			Title:       movies[i].Title,
			ReleaseYear: movies[i].ReleaseYear,
			Quality:     movies[i].Quality,
			Id:          movies[i].Id,
		})
	}
	return final_movies, nil
}
