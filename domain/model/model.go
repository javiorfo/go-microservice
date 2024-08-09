package model

import "github.com/javiorfo/go-microservice/internal/auditory"

type Dummy struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Info string `json:"info"`
	auditory.Auditable
}
