package Models

import "gorm.io/gorm"

type Role struct {
	ID   uint32 `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name string `gorm:"size:200;column:name" json:"name"`
}

func (role Role) TableName() string {
	return "roles"
}

func MigrateRoles(db *gorm.DB) error {
	err := db.AutoMigrate(&Role{})
	return err
}
