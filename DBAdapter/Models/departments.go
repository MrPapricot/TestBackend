package Models

import (
	"gorm.io/gorm"
)

type Department struct {
	ID                  uint64            `gorm:"primaryKey;type:bigInt;autoIncrement;column:id" json:"id"`
	Name                string            `gorm:"size:200;column:name;uniqueIndex" json:"name"`
	EngineeringSchoolId uint32            `gorm:"column:engineering_school_id;not null;index" json:"engineering_school_id"`
	EngineeringSchool   EngineeringSchool `gorm:"foreignKey:EngineeringSchoolId;references:ID" json:"engineering_school"`
}

func (department Department) TableName() string {
	return "departments"
}

func MigrateDepartments(db *gorm.DB) error {
	return db.AutoMigrate(&Department{})
}
