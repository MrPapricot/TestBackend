package Models

import (
	"gorm.io/gorm"
)

type Roadmap struct {
	ID               uint64         `gorm:"primaryKey;column:id;autoIncrement"`
	SpecializationID uint64         `gorm:"not null;index;column:specialization_id"`
	Specialization   Specialization `gorm:"foreignKey:SpecializationID;references:ID"`
}

func (roadmap Roadmap) TableName() string {
	return "roadmaps"
}

func MigrateRoadmaps(db *gorm.DB) error {
	return db.AutoMigrate(&Roadmap{})
}
