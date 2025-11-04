package Models

import (
	"gorm.io/gorm"
)

type Specialization struct {
	ID                  uint64            `gorm:"primaryKey;column:id"`
	Name                string            `gorm:"type:varchar(200);column:name"`
	TrainingDirectionID uint64            `gorm:"not null;index;column:training_direction_id"`
	TrainingDirection   TrainingDirection `gorm:"foreignKey:TrainingDirectionID;references:ID"`
}

func (specialization Specialization) TableName() string {
	return "specializations"
}

func MigrateSpecializations(db *gorm.DB) error {
	return db.AutoMigrate(&Specialization{})
}
