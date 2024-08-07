package persistence

type DummyEntity struct {
    ID   uint `gorm:"primaryKey;autoIncrement"`
	Info string
}

func (DummyEntity) TableName() string {
  return "dummies"
}
