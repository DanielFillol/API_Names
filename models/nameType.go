package models

import (
	"gorm.io/gorm"
	"time"
)

//NameType main struct
type NameType struct {
	gorm.Model
	Name           string `json:"Name,omitempty"`
	Classification string `json:"Classification,omitempty"`
	Metaphone      string `json:"Metaphone,omitempty"`
	NameVariations string `json:"NameVariations,omitempty"`
}

//NameVar struct for organizing name variations by Levenshtein
type NameVar struct {
	Name        string
	Levenshtein float32
}

//MetaphoneR only use for SearchSimilarNames return
type MetaphoneR struct {
	ID             uint           `json:"ID,omitempty"`
	CreatedAt      time.Time      `json:"CreatedAt,omitempty"`
	UpdatedAt      time.Time      `json:"UpdatedAt,omitempty"`
	DeletedAt      gorm.DeletedAt `json:"DeletedAt,omitempty"`
	Name           string         `json:"Name,omitempty"`
	Classification string         `json:"Classification,omitempty"`
	Metaphone      string         `json:"Metaphone,omitempty"`
	NameVariations []string       `json:"NameVariations,omitempty"`
}
