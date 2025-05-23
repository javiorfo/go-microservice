package model

import (
	"github.com/javiorfo/go-microservice-lib/auditory"
	"github.com/javiorfo/go-microservice-lib/validation"
)

// Dummy represents a dada structure
type Dummy struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Info   string `json:"info"`
	Status Status `json:"status"`
	auditory.Auditable
}

type Status = string

const (
	enable   Status = "ON"
	disabled Status = "OFF"
)

var ValidateStatus = validation.NewEnumValidator("status", "status", enable, disabled)
