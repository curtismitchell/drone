package builtin

// DO NOT EDIT
// code generated by go:generate

import (
	"database/sql"
	"encoding/json"

	. "github.com/drone/drone/pkg/types"
)

var _ = json.Marshal

// generic database interface, matching both *sql.Db and *sql.Tx
type userDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

func getUser(db userDB, query string, args ...interface{}) (*User, error) {
	row := db.QueryRow(query, args...)
	return scanUser(row)
}

func getUsers(db userDB, query string, args ...interface{}) ([]*User, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanUsers(rows)
}

func createUser(db userDB, query string, v *User) error {
	var v0 string
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 bool
	var v6 bool
	var v7 string
	v0 = v.Login
	v1 = v.Token
	v2 = v.Secret
	v3 = v.Email
	v4 = v.Avatar
	v5 = v.Active
	v6 = v.Admin
	v7 = v.Hash

	res, err := db.Exec(query,
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
	)
	if err != nil {
		return err
	}

	v.ID, err = res.LastInsertId()
	return err
}

func updateUser(db userDB, query string, v *User) error {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 bool
	var v7 bool
	var v8 string
	v0 = v.ID
	v1 = v.Login
	v2 = v.Token
	v3 = v.Secret
	v4 = v.Email
	v5 = v.Avatar
	v6 = v.Active
	v7 = v.Admin
	v8 = v.Hash

	_, err := db.Exec(query,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
		&v0,
	)
	return err
}

func scanUser(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string
	var v3 string
	var v4 string
	var v5 string
	var v6 bool
	var v7 bool
	var v8 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
		&v3,
		&v4,
		&v5,
		&v6,
		&v7,
		&v8,
	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.ID = v0
	v.Login = v1
	v.Token = v2
	v.Secret = v3
	v.Email = v4
	v.Avatar = v5
	v.Active = v6
	v.Admin = v7
	v.Hash = v8

	return v, nil
}

func scanUsers(rows *sql.Rows) ([]*User, error) {
	var err error
	var vv []*User
	for rows.Next() {
		var v0 int64
		var v1 string
		var v2 string
		var v3 string
		var v4 string
		var v5 string
		var v6 bool
		var v7 bool
		var v8 string
		err = rows.Scan(
			&v0,
			&v1,
			&v2,
			&v3,
			&v4,
			&v5,
			&v6,
			&v7,
			&v8,
		)
		if err != nil {
			return vv, err
		}

		v := &User{}
		v.ID = v0
		v.Login = v1
		v.Token = v2
		v.Secret = v3
		v.Email = v4
		v.Avatar = v5
		v.Active = v6
		v.Admin = v7
		v.Hash = v8
		vv = append(vv, v)
	}
	return vv, rows.Err()
}

const stmtUserSelectList = `
SELECT
 user_id
,user_login
,user_token
,user_secret
,user_email
,user_avatar
,user_active
,user_admin
,user_hash
FROM users
`

const stmtUserSelectRange = `
SELECT
 user_id
,user_login
,user_token
,user_secret
,user_email
,user_avatar
,user_active
,user_admin
,user_hash
FROM users
LIMIT ? OFFSET ?
`

const stmtUserSelect = `
SELECT
 user_id
,user_login
,user_token
,user_secret
,user_email
,user_avatar
,user_active
,user_admin
,user_hash
FROM users
WHERE user_id = ?
`

const stmtUserSelectUserLogin = `
SELECT
 user_id
,user_login
,user_token
,user_secret
,user_email
,user_avatar
,user_active
,user_admin
,user_hash
FROM users
WHERE user_login = ?
`

const stmtUserSelectCount = `
SELECT count(1)
FROM users
`

const stmtUserInsert = `
INSERT INTO users (
 user_login
,user_token
,user_secret
,user_email
,user_avatar
,user_active
,user_admin
,user_hash
) VALUES (?,?,?,?,?,?,?,?);
`

const stmtUserUpdate = `
UPDATE users SET
 user_login = ?
,user_token = ?
,user_secret = ?
,user_email = ?
,user_avatar = ?
,user_active = ?
,user_admin = ?
,user_hash = ?
WHERE user_id = ?
`

const stmtUserDelete = `
DELETE FROM users
WHERE user_id = ?
`

const stmtUserTable = `
CREATE TABLE IF NOT EXISTS users (
 user_id	INTEGER PRIMARY KEY AUTOINCREMENT
,user_login	VARCHAR
,user_token	VARCHAR
,user_secret	VARCHAR
,user_email	VARCHAR
,user_avatar	VARCHAR
,user_active	BOOLEAN
,user_admin	BOOLEAN
,user_hash	VARCHAR
);
`

const stmtUserUserLoginIndex = `
CREATE UNIQUE INDEX IF NOT EXISTS ux_user_login ON users (user_login);
`