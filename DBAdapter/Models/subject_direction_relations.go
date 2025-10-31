package Models

import (
	"gorm.io/gorm"
)

type SubjectDirectionRelation struct {
	SubjectID           uint64            `gorm:"not null;index;column:subject_id"`
	Subject             Subject           `gorm:"foreignKey:SubjectID;references:ID"`
	TrainingDirectionID uint64            `gorm:"not null;index;column:training_direction_id"`
	TrainingDirection   TrainingDirection `gorm:"foreignKey:TrainingDirectionID;references:ID"`
}

func (relation SubjectDirectionRelation) TableName() string {
	return "subject_direction_relations"
}

func MigrateSubjectDirectionRelations(db *gorm.DB) error {
	return db.AutoMigrate(&SubjectDirectionRelation{})
}
