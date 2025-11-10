package Models

import (
	"gorm.io/gorm"
)

type RelationType struct {
	ID   uint16 `gorm:"type:smallInt;primaryKey;column:id;autoIncrement"`
	Code string `gorm:"uniqueIndex;column:code;type:varchar(20)"`
	Name string `gorm:"uniqueIndex;type:varchar(200);column:name"`
}

func (relationType RelationType) TableName() string {
	return "relation_types"
}

func MigrateRelationTypes(db *gorm.DB) error {
	return db.AutoMigrate(&RelationType{})
}
