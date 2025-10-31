package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseTpuUser struct {
	UUID       uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:uuid" json:"uuid"`
	FirstName  string    `gorm:"column:first_name;type:varchar(100)" json:"first_name"`
	LastName   string    `gorm:"column:last_name;type:varchar(100)" json:"last_name"`
	MiddleName string    `gorm:"column:middle_name;type:varchar(100)" json:"middle_name"`
	Login      string    `gorm:"uniqueIndex;column:login;type:varchar(10)" json:"login"`
}

func (user BaseTpuUser) TableName() string {
	return "base_tpu_users"
}

func MigrateBaseTpuUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&BaseTpuUser{})
	return err
}
