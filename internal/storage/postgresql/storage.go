package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"medodsAuth/internal/models"
)

var DB *Storage

type Storage struct {
	*pgx.Conn
}

func ConnectDB(storagePath string) {
	const op = "storage.postgresql.New"

	conn, err := pgx.Connect(context.Background(), storagePath)
	if err != nil {
		log.Fatalf("%s : %v", op, err)
	}

	log.Println("storage successfully configured")
	DB = &Storage{Conn: conn}
}

func (db *Storage) Close() {
	db.Conn.Close(context.Background())
}

func (db *Storage) SaveUser(user *models.User) error {
	const (
		op    = "storage.postgresql.SaveUser"
		query = "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4);"
	)

	_, err := db.Exec(context.Background(), query, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Printf("%s : %v", op, err)
		return fmt.Errorf("query can't be executed : %v", err.Error())
	}

	return nil
}
