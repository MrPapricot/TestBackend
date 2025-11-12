package Models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type RoadmapNode struct {
	ID        uint64         `gorm:"primaryKey;column:id;autoIncrement"`
	JSONData  datatypes.JSON `gorm:"type:jsonb;column:jsondata"`
	RoadmapID uint64         `gorm:"not null;index;column:roadmap_id"`
	Roadmap   Roadmap        `gorm:"foreignKey:RoadmapID;references:ID;constraint:OnDelete:CASCADE"`
}

func (node RoadmapNode) TableName() string {
	return "roadmap_nodes"
}

func MigrateRoadmapNodes(db *gorm.DB) error {
	return db.AutoMigrate(&RoadmapNode{})
}
