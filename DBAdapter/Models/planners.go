package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Planner struct {
	UUID         uuid.UUID   `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:uuid" json:"id"`
	BaseUserUUID uuid.UUID   `gorm:"type:uuid;column:base_user_uuid;not null;index" json:"base_user_uuid"`
	BaseUser     BaseTpuUser `gorm:"foreignKey:BaseUserUUID;references:UUID" json:"base_user"`
}

func (planner Planner) TableName() string {
	return "planners"
}

func MigratePlanners(db *gorm.DB) error {
	return db.AutoMigrate(&Planner{})
}
