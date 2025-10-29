package DBConnection

import (
	"database/sql"
	"fmt"
)

type ConnectionInfo struct {
	Host     string
	Port     int
	DBName   string
	User     string
	Password string
}

func InitConnection(data ConnectionInfo) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", data.Host, data.Port, data.User, data.Password, data.DBName)
	conn, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к серверу. Проверьте доступен ли он")
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
