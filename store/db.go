package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gocarbe/shared"
)

// TODO : ALL TEMP

var store *sql.DB

func SetUpDb() {
	db, err := sql.Open("sqlite3", "./store.db")
	store = db
	if err != nil {
		shared.LOG(err.Error())
	}

	sqlStmt := `
	create table if not exists emails (id integer not null primary key, email text unique);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		shared.LOG(err.Error())
		return
	}
}

func CloseDb() {
	store.Close()
}

func AddEmail(email string) {
	tx, err := store.Begin()
	if err != nil {
		shared.LOG(err.Error())
	}
	stmt, err := tx.Prepare("insert into emails(email) values(?)")
	if err != nil {
		shared.LOG(err.Error())
	}

	_, err = stmt.Exec(email)
	if err != nil {
		shared.LOG(err.Error())
	}

	err = stmt.Close()
	if err != nil {
		shared.LOG(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		shared.LOG(err.Error())
	}
}

func GetEmails() []string {
	emails := make([]string, 0)

	rows, err := store.Query("select email from emails")
	if err != nil {
		shared.LOG(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		err = rows.Scan(&email)

		emails = append(emails, email)

		if err != nil {
			shared.LOG(err.Error())
		}
	}
	err = rows.Err()
	if err != nil {
		shared.LOG(err.Error())
	}

	return emails
}
