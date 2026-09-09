package mysqlcast

import (
	"database/sql"
	"movieapp/api/param/repositoryparam"
	"movieapp/entity"

	"github.com/jmoiron/sqlx"
)

func (db *DB) GetAllCasts() ([]entity.Cast, error) {
	conn := db.conn.Conn()

	casts := make([]repositoryparam.Cast, 0)
	err := conn.Select(&casts, `
		SELECT id, name
		FROM cast
		ORDER BY id ASC
	`)
	if err != nil {
		return nil, err
	}

	final_casts := make([]entity.Cast, 0)
	for i := range casts {
		final_casts = append(final_casts, entity.Cast{
			Name: casts[i].Name,
			Id:   casts[i].Id,
		})
	}

	return final_casts, nil
}
func (db *DB) GetCastById(id int) (*entity.Cast, error) {
	conn := db.conn.Conn()

	var c repositoryparam.Cast

	err := conn.Get(&c, `
		SELECT id, name
		FROM cast
		WHERE id=?
	`, id)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &entity.Cast{
		Name: c.Name,
		Id:   c.Id,
	}, nil
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

	err := conn.Get(&exists, `
		SELECT EXISTS(
			SELECT 1 FROM cast WHERE id=?
		)
	`, id)

	if err != nil {
		return nil, err
	}

	return &exists, nil
}

func (db *DB) GetCastByIds(ids []int) ([]entity.Cast, error) {
	conn := db.conn.Conn()

	q, args, err := sqlx.In(`
	SELECT id,name FROM cast
	WHERE id IN (?)
	ORDER BY id ASC
	`, ids)

	if err != nil {
		return nil, err
	}

	casts := make([]repositoryparam.Cast, 0)
	err = conn.Select(&casts, q, args...)

	if err != nil {
		return nil, err
	}

	final_casts := make([]entity.Cast, 0)
	for i := range casts {
		final_casts = append(final_casts, entity.Cast{
			Name: casts[i].Name,
			Id:   casts[i].Id,
		})
	}

	return final_casts, nil
}
