package storage

import (
	"context"
	"database/sql"
	"errors"
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
	const op = "storage.postgresql.ConnectDB"

	conn, err := pgx.Connect(context.Background(), storagePath)
	if err != nil {
		log.Fatalf("%s : %v", op, err)
	}

	log.Println("storage successfully configured")
	DB = &Storage{Conn: conn}
}

func (db *Storage) Close() {
	db.Conn.Close(context.Background())
	log.Println("storage connection closed gracefully")
}

func (db *Storage) SaveUser(user *models.User) error {
	const (
		op    = "storage.postgresql.SaveUser"
		query = "INSERT INTO users (guid ,first_name, last_name, email, password) VALUES ($1, $2, $3, $4, $5);"
	)

	_, err := db.Exec(context.Background(), query, user.GUID, user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		log.Printf("%s : %v", op, err)
		return fmt.Errorf("query can't be executed : %v", err.Error())
	}

	return nil
}

func (db *Storage) SaveToken(token *models.RefreshToken) error {
	const (
		op    = "storage.postgresql.SaveToken"
		query = "INSERT INTO refresh_tokens (user_guid, ip_address, refresh_token_hash, token_id) VALUES ($1, $2, $3, $4);"
	)

	_, err := db.Exec(context.Background(), query, token.GUID, token.IP, token.RefreshTokenHash, token.ID)
	if err != nil {
		log.Printf("%s : %v", op, err)
		return fmt.Errorf("query can't be executed : %v", err.Error())
	}

	return nil
}

func (db *Storage) GetUserByGUID(guid string) (*models.User, error) {
	const (
		op    = "storage.postgresql.GetUserByGUID"
		query = "SELECT email, password, last_name, first_name FROM users WHERE guid=$1"
	)
	var user = &models.User{}
	row := db.QueryRow(context.Background(), query, guid)

	if err := row.Scan(&user.Email, &user.Password, &user.LastName, &user.FirstName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s : user with GUID %s not found", op, guid)
		}
		log.Printf("%s : %s\n", op, err.Error())
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return user, nil
}

func (db *Storage) GetRefreshToken(guid, tokenID string) (*models.RefreshToken, error) {
	const (
		op    = "storage.postgresql.GetRefreshToken"
		query = "SELECT refresh_token_hash FROM refresh_tokens WHERE user_guid=$1 AND token_id=$2"
	)
	var refreshToken models.RefreshToken

	err := db.QueryRow(context.Background(), query, guid, tokenID).Scan(&refreshToken.RefreshTokenHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s : %s", op, "token not found")
		}
		return nil, fmt.Errorf("%s : %w", op, err)
	}

	return &refreshToken, nil
}

func (db *Storage) UpdateRefreshToken(guid, tokenID, newRefreshTokenHash string) error {
	const (
		op          = "storage.postgresql.UpdateRefreshToken"
		query       = "UPDATE refresh_tokens SET refresh_token_hash=$1 WHERE user_guid=$2 AND token_id=$3"
		deleteQuery = "DELETE FROM refresh_tokens WHERE user_guid=$1 AND token_id!=$2"
	)

	_, err := db.Exec(context.Background(), query, newRefreshTokenHash, guid, tokenID)
	if err != nil {
		return fmt.Errorf("%s : %v", op, err)
	}

	_, err = db.Exec(context.Background(), deleteQuery, guid, tokenID)
	if err != nil {
		return fmt.Errorf("%s : %v", op, err)
	}

	return nil
}
