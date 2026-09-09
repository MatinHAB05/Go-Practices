package mysqlmovie

import (
	"database/sql"
	"movieapp/entity"
)

func (db *DB) GetAllMovies() ([]entity.Movie, error) {
	conn := db.conn.Conn()

	rows, err := conn.Query("SELECT id,title,quality,release_year FROM movie")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	movies := make([]entity.Movie, 0)

	for rows.Next() {
		var m entity.Movie

		err := rows.Scan(&m.Id, &m.Title, &m.Quality, &m.ReleaseYear)
		if err != nil {
			return nil, err
		}

		movies = append(movies, m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func (db *DB) GetMovieById(id int) (*entity.Movie, error) {
	conn := db.conn.Conn()

	row := conn.QueryRow(`
		SELECT id, title, quality, release_year
		FROM movie
		WHERE id=?
	`, id)

	var m entity.Movie

	err := row.Scan(
		&m.Id,
		&m.Title,
		&m.Quality,
		&m.ReleaseYear,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &m, nil
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

	err := conn.QueryRow(`
		SELECT EXISTS(
			SELECT id FROM movie WHERE id=?
		)
	`, id).Scan(&exists)

	if err != nil {
		return nil, err
	}

	return &exists, nil
}
