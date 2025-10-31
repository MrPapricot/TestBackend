package Models

import "gorm.io/gorm"

type QualificationLevel struct {
	ID   uint32 `gorm:"primaryKey;type:smallInt;autoIncrement;column:id" json:"id"`
	Name string `gorm:"type:varchar(200);column:name;uniqueIndex" json:"name"`
}

func (level QualificationLevel) TableName() string {
	return "qualification_levels"
}

func MigrateQualificationLevels(db *gorm.DB) error {
	return db.AutoMigrate(&QualificationLevel{})
}
