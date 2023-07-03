package models

import "gorm.io/gorm"

type Shard struct {
	gorm.Model
	Name         string        `gorm:"type:text;not null" json:"name"`
	Vessel       string        `gorm:"type:text;not null" json:"vessel"`
	Planet       string        `gorm:"type:text;not null" json:"planet"`
	MagicSystems []MagicSystem `json:"magic_systems"`
}

type ShardReturn struct {
	ID           uint                `json:"id"`
	Name         string              `json:"name"`
	Vessel       string              `json:"vessel"`
	Planet       string              `json:"planet"`
	MagicSystems []MagicSystemReturn `json:"magic_systems" gorm:"foreignKey:ShardID"`
}
