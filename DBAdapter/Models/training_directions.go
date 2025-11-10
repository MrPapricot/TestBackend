package Models

import (
	"gorm.io/gorm"
)

type TrainingDirection struct {
	ID                   uint64             `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name                 string             `gorm:"type:varchar(200);column:name" json:"name"`
	Duration             uint16             `gorm:"column:duration" json:"duration"`
	QualificationLevelID uint16             `gorm:"type:smallInt;not null;index;column:qualification_level_id" json:"qualification_level_id"`
	QualificationLevel   QualificationLevel `gorm:"foreignKey:QualificationLevelID;references:ID"`
}

func (training_directions TrainingDirection) TableName() string {
	return "training_directions"
}

func MigrateTrainingDirections(db *gorm.DB) error {
	return db.AutoMigrate(&TrainingDirection{})
}
