package Models

import "gorm.io/gorm"

type EngineeringSchool struct {
	ID   uint32 `gorm:"primaryKey;type:integer;autoIncrement;column:id" json:"id"`
	Name string `gorm:"type:varchar(200);column:name;uinqueIndex" json:"name"`
}

func (school EngineeringSchool) TableName() string {
	return "engineering_schools"
}

func MigrateEngineeringSchools(db *gorm.DB) error {
	return db.AutoMigrate(&EngineeringSchool{})
}
