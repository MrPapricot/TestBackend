package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseTpuUser struct {
	UUID       uuid.UUID `gorm:"primaryKey;default:gen_random_uuid();column:uuid" json:"uuid"`
	FirstName  string    `grom:"column:first_name;size:100" json:"first_name"`
	LastName   string    `grom:"column:last_name;size:100" json:"last_name"`
	MiddleName string    `grom:"column:middle_name;size:100" json:"middle_name"`
	Login      string    `gorm:"uniqueIndex;column:login" json:"login"`
}

func (user BaseTpuUser) TableName() string {
	return "base_tpu_users"
}

func MigrateBaseTpuUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&BaseTpuUser{})
	return err
}
