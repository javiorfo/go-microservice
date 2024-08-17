package model

import "github.com/javiorfo/go-microservice-lib/auditory"

// Dummy represents a dada structure
type Dummy struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Info string `json:"info"`
	auditory.Auditable
}
