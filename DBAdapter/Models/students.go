package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	UUID            uuid.UUID     `gorm:"primaryKey;column:uuid" json:"uuid"`
	BaseUserUUID    uuid.UUID     `gorm:"not null;index;column:base_user_uuid" json:"base_user_id"`
	BaseUser        BaseTpuUser   `gorm:"foreignKey:BaseUserUUID;references:UUID"`
	StudentIDNumber string        `gorm:"size:200;column:student_id_number" json:"student_id_number"`
	EducationTypeID uint32        `gorm:"not null;index;column:education_type_id" json:"education_type_id"`
	EducationType   EducationType `gorm:"foreignKey:EducationTypeID;references:ID"`
	GroupID         uint64        `gorm:"not null;index;column:group_id" json:"group_id"`
	Group           Group         `gorm:"foreignKey:GroupID;references:ID"`
}

func (student Student) TableName() string {
	return "students"
}

func MigrateStudents(db *gorm.DB) error {
	return db.AutoMigrate(&Student{})
}
