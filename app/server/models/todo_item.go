package models

import (
	"database/sql"
)

type TodoItemDb struct {
	MySqlDB *sql.DB
}

func (t *TodoItemDb) Add(id int, login string, uid string) (int, error) {
	stmt := `INSERT INTO auth (id, login, uid, created, expires)
    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := t.MySqlDB.Exec(stmt, id, login, uid)
	if err != nil {
		panic(err)
	}
	lastId, errInsert := result.LastInsertId()
	if errInsert != nil {
		return 0, errInsert
	}

	return int(lastId), nil
}
