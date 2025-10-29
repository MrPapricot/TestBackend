package DBAdapter

import (
	"backend/DBConnection"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type DBAdapter struct {
	connection *sql.DB
}

type BaseTpuUser struct {
	UUID       uuid.UUID
	Name       string
	LastName   string
	MiddleName string
	Login      string
}

func InitAdapter(conn_info DBConnection.ConnectionInfo) (*DBAdapter, error) {
	conn, err := DBConnection.InitConnection(conn_info)
	if err != nil {
		return nil, err
	}
	return &DBAdapter{
		connection: conn,
	}, nil
}

func (adapter *DBAdapter) GetBaseUserByUUID(uuid uuid.UUID) (*BaseTpuUser, error) {
	query := `
        SELECT uuid, name, last_name, middle_name, login 
        FROM base_tpu_users 
        WHERE uuid = $1
    `

	var base BaseTpuUser
	err := adapter.connection.QueryRow(query, uuid).Scan(
		&base.UUID,
		&base.Name,
		&base.LastName,
		&base.MiddleName,
		&base.Login,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("database error: %w", err)
	}

	return &base, nil
}

func (adapter *DBAdapter) GetAllBaseUsers() ([]BaseTpuUser, error) {
	query := "SELECT * FROM base_tpu_users"
	var (
		users     []BaseTpuUser
		next_user BaseTpuUser
	)
	res, err := adapter.connection.Query(query)
	if err != nil {
		return nil, err
	}
	for res.Next() {
		res.Scan(
			&next_user.UUID,
			&next_user.Name,
			&next_user.LastName,
			&next_user.MiddleName,
			&next_user.Login,
		)
		users = append(users, next_user)
	}
	return users, nil
}
