package repositories

import (
	"database/sql"
	"log"

	"wahyuade.com/simple-job-api/helpers"
	"wahyuade.com/simple-job-api/models"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type UserRespository struct {
	connection *sql.DB
}

func (uR UserRespository) GetByUsername(username string) (user models.User) {
	stmt, err := uR.connection.Prepare(`SELECT id, name, username, password FROM "user" WHERE username = ?`)
	if err != nil {
		log.Println(err)
	}
	stmt.QueryRow(username).Scan(&user.Id, &user.Name, &user.Username, &user.Password)
	return user
}

func (uR UserRespository) Insert(user models.User) bool {
	password, _ := helpers.Hash(user.Password)
	stmt, err := uR.connection.Prepare(`INSERT INTO "user" (id, name, username, password) VALUES (?,?,?,?)`)
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(
		uuid.NewString(),
		user.Name,
		user.Username,
		password,
	)
	return err == nil
}

func (uR UserRespository) IsUsernameExist(username string) bool {
	stmt, err := uR.connection.Prepare(`SELECT username FROM "user" WHERE username = ?`)
	if err != nil {
		log.Println(err)
	}
	var uname string
	stmt.QueryRow(username).Scan(&uname)
	return uname == username
}

func newUserRepository() UserRespository {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}
	db.Exec(`CREATE TABLE IF NOT EXISTS "user" (
		"id" text PRIMARY KEY,
		"name" text,
		"username" text UNIQUE,
		"password" text
	)`)
	return UserRespository{
		connection: db,
	}
}
