package mappers

import (
	"github.com/javiorfo/go-microservice/adapter/persistence"
	"github.com/javiorfo/go-microservice/domain/model"
)

func DummyEntityToDummy(de persistence.DummyEntity) model.Dummy {
	return model.Dummy{
        Info: de.Info,
	}
}

func DummyToDummyEntity(d model.Dummy) persistence.DummyEntity {
	return persistence.DummyEntity{
        Info: d.Info,
	}
}
