package Models

import (
	"gorm.io/gorm"
)

type SubjectTutorRelation struct {
	SubjectID   uint64  `gorm:"not null;index;column:subject_id"`
	Subject     Subject `gorm:"foreignKey:SubjectID;references:ID"`
	TutorRoleID uint32  `gorm:"not null;index;column:tutor_role_id"`
	TutorRole   Role    `gorm:"foreignKey:TutorRoleID;references:ID"`
}

func (relation SubjectTutorRelation) TableName() string {
	return "subject_tutor_relations"
}

func MigrateSubjectTutorRelations(db *gorm.DB) error {
	return db.AutoMigrate(&SubjectTutorRelation{})
}
