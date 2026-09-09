package mysqlcast

import (
	"database/sql"
	"fmt"
	"movieapp/entity"
	"strings"
)

func (db *DB) GetAllCasts() ([]entity.Cast, error) {
	conn := db.conn.Conn()

	rows, err := conn.Query(`
		SELECT id, name
		FROM cast
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	casts := make([]entity.Cast, 0)

	for rows.Next() {
		var c entity.Cast

		err := rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, err
		}

		casts = append(casts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return casts, nil
}
func (db *DB) GetCastById(id int) (*entity.Cast, error) {
	conn := db.conn.Conn()

	row := conn.QueryRow(`
		SELECT id, name
		FROM cast
		WHERE id=?
	`, id)

	var c entity.Cast

	err := row.Scan(
		&c.Id,
		&c.Name,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (db *DB) Create(c entity.Cast) (*entity.Cast, error) {
	conn := db.conn.Conn()

	res, err := conn.Exec(`
		INSERT INTO cast (name)
		VALUES (?)
	`, c.Name)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	c.Id = int(id)

	return &c, nil
}

func (db *DB) Delete(id int) (*entity.Cast, error) {
	conn := db.conn.Conn()

	cast, err := db.GetCastById(id)
	if err != nil {
		return nil, err
	}

	if cast == nil {
		return nil, sql.ErrNoRows
	}

	_, err = conn.Exec(`
		DELETE FROM cast
		WHERE id=?
	`, id)

	if err != nil {
		return nil, err
	}

	return cast, nil
}

func (db *DB) IsExist(id int) (*bool, error) {
	conn := db.conn.Conn()

	var exists bool

	err := conn.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM cast WHERE id=?
		)
	`, id).Scan(&exists)

	if err != nil {
		return nil, err
	}

	return &exists, nil
}

func (db *DB) GetCastByIds(ids []int) ([]entity.Cast, error) {
	conn := db.conn.Conn()

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))

	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	query := fmt.Sprintf(`
	SELECT id,name FROM cast
	WHERE id IN (%s)
	ORDER BY id ASC
	`, strings.Join(placeholders, ","))

	rows, err := conn.Query(query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	casts := make([]entity.Cast, 0)

	for rows.Next() {
		var c entity.Cast
		err := rows.Scan(&c.Id, &c.Name)
		if err != nil {
			return nil, err
		}
		casts = append(casts, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return casts, nil
}
