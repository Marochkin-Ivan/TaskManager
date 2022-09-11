package db

import (
	"log"
	"taskmanager/internal/models"
)

func (d *DB) RegistrationUser(u models.User) (string, bool, error) {
	stmt := `select uid from users where login=$1`
	rows, err := d.Conn.Query(stmt, u.Login)
	if err != nil {
		return "", false, err
	}
	if rows.Err() != nil {
		return "", false, rows.Err()
	}

	if rows.Next() {
		return "", true, nil
	}
	rows.Close()

	uid := GenerateUID()
	stmt = `insert into users(uid,login,pass) values($1,$2,$3)`
	_, err = d.Conn.Exec(stmt, uid, u.Login, u.Password)
	if err != nil {
		return "", false, err
	}

	return uid, false, nil
}

func (d *DB) AuthorizationUser(u models.User) (string, bool, error) {
	stmt := `select uid from users where login=$1 and pass=$2`
	rows, err := d.Conn.Query(stmt, u.Login, u.Password)
	if err != nil {
		return "", false, err
	}
	if rows.Err() != nil {
		return "", false, rows.Err()
	}

	defer rows.Close()
	var uid string
	if rows.Next() {
		err = rows.Scan(&uid)
		if err != nil {
			log.Println(err)
			return "", false, err
		}
	} else {
		return "", false, nil
	}

	return uid, true, nil
}
