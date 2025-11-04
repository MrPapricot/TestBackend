package Models

import (
	"gorm.io/gorm"
)

type SubjectSpecializationRelation struct {
	SubjectID        uint64         `gorm:"not null;index;column:subject_id"`
	Subject          Subject        `gorm:"foreignKey:SubjectID;references:ID"`
	SpecializationID uint64         `gorm:"not null;index;column:specialization_id"`
	Specialization   Specialization `gorm:"foreignKey:SpecializationID;references:ID"`
	RelationTypeID   uint64         `gorm:"not null;index;column:relation_type_id"`
	RelationType     RelationType   `gorm:"foreignKey:RelationTypeID;references:ID"`
}

func (relation SubjectSpecializationRelation) TableName() string {
	return "subject_specialization_relations"
}

func MigrateSubjectSpecializationRelations(db *gorm.DB) error {
	return db.AutoMigrate(&SubjectSpecializationRelation{})
}
