package db

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type UsersDb struct {
	MySqlDB *sql.DB
}

//func (t *TodoItemDb) Add(login string) (int, error) {
//	stmt := `INSERT INTO user (id, login, uid, created, expires)
//    VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
//
//	result, err := t.MySqlDB.Exec(stmt, login)
//	if err != nil {
//		panic(err)
//	}
//	lastId, errInsert := result.LastInsertId()
//	if errInsert != nil {
//		return 0, errInsert
//	}
//
//	return int(lastId), nil
//}

func (t *UsersDb) IsUserExist(login string) error {
	var id int
	err := t.MySqlDB.QueryRow(`SELECT id FROM users WHERE login = ?`, login).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("Нет такого пользователя!")
		}
		log.Info("Ошибка")
	}

	return nil
}
