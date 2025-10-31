package Models

import (
	"gorm.io/gorm"
)

type Group struct {
	ID                  uint64            `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	TrainingDirectionID uint64            `gorm:"not null;index;column:training_direction_id" json:"training_direction_id"`
	TrainingDirection   TrainingDirection `gorm:"foreignKey:TrainingDirectionID;references:ID"`
	Name                string            `gorm:"type:varchar(20);column:name"`
}

func (group Group) TableName() string {
	return "groups"
}

func MigrateGroups(db *gorm.DB) error {
	return db.AutoMigrate(&Group{})
}
