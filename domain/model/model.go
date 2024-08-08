package model

type Dummy struct {
    ID   uint   `json:"-" gorm:"primaryKey;autoIncrement"`
	Info string `json:"info"`
}
