package Models

import (
	"gorm.io/gorm"
)

type Subject struct {
	ID   uint64 `gorm:"primaryKey;type:bigInt;autoIncrement;column:id" json:"id"`
	Name string `gorm:"type:varchar(200);column:name;index" json:"name"`
}

func (subject Subject) TableName() string {
	return "subjects"
}

func MigrateSubjects(db *gorm.DB) error {
	return db.AutoMigrate(&Subject{})
}
