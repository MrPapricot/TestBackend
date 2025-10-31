package Models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tutor struct {
	UUID               uuid.UUID   `gorm:"primaryKey;type:uuid;default:gen_random_uuid();column:uuid" json:"id"`
	BaseUserUUID       uuid.UUID   `gorm:"type:uuid;column:base_user_uuid;not null;index" json:"base_user_uuid"`
	BaseUser           BaseTpuUser `gorm:"foreignKey:BaseUserUUID;references:UUID" json:"base_user"`
	PersonalWebsiteUrl string      `gorm:"type:varchar(350);column:personal_website_url" json:"personal_website_url"`
}

func (tutor Tutor) TableName() string {
	return "tutors"
}

func MigrateTutors(db *gorm.DB) error {
	return db.AutoMigrate(&Tutor{})
}
