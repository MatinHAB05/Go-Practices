package mysqlmoviecast

func (db *DB) Create(movie_id int, cast_id int) (*int, *int, error) {
	conn := db.conn.Conn()
	_, err := conn.Exec(`
	INSERT INTO moviecast VALUES
	(?,?)
	`, movie_id, cast_id)

	if err != nil {
		return nil, nil, err
	}

	return &movie_id, &cast_id, nil
}

func (db *DB) IsExist(movie_id int, cast_id int) (*bool, error) {
	conn := db.conn.Conn()

	row := conn.QueryRow(`
		SELECT EXISTS(
			SELECT 1
			FROM moviecast
			WHERE movie_id = ? AND cast_id = ?
		)
	`, movie_id, cast_id)

	var exist bool
	err := row.Scan(&exist)
	if err != nil {
		return nil, err
	}

	return &exist, nil
}
