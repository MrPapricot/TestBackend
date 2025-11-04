package Models

import (
	"gorm.io/gorm"
)

type RelationType struct {
	ID   uint64 `gorm:"primaryKey;column:id"`
	Name string `gorm:"type:varchar(100);column:name"`
}

func (relationType RelationType) TableName() string {
	return "relation_types"
}

func MigrateRelationTypes(db *gorm.DB) error {
	return db.AutoMigrate(&RelationType{})
}
