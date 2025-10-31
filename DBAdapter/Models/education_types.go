package Models

import "gorm.io/gorm"

type EducationType struct {
	ID   uint16 `gorm:"primaryKey;type:smallInt;autoIncrement;column:id" json:"id"`
	Name string `gorm:"size:200;column:name;uniqueIndex" json:"name"`
}

func (education_type EducationType) TableName() string {
	return "education_types"
}

func MigrateEducationTypes(db *gorm.DB) error {
	return db.AutoMigrate(&EducationType{})
}
