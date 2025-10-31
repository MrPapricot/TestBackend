package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TutorDepartmentRelation struct {
	TutorUUID    uuid.UUID  `gorm:"not null;index;column:tutor_uuid"`
	Tutor        Tutor      `gorm:"foreignKey:TutorUUID;references:UUID"`
	DepartmentID uint64     `gorm:"not null;index;column:department_id"`
	Department   Department `gorm:"foreignKey:DepartmentID;references:ID"`
	TutorRoleID  uint32     `gorm:"not null;index;column:tutor_role_id"`
	TutorRole    Role       `gorm:"foreignKey:TutorRoleID;references:ID"`
}

func (relation TutorDepartmentRelation) TableName() string {
	return "tutor_department_relations"
}

func MigrateTutorDepartmentRelations(db *gorm.DB) error {
	return db.AutoMigrate(&TutorDepartmentRelation{})
}
