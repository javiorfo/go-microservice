package persistence

import "github.com/javiorfo/go-microservice/domain/model"

func DummyEntityToDummy(de DummyEntity) model.Dummy {
	return model.Dummy{
        Info: de.Info,
	}
}

func DummyToDummyEntity(de model.Dummy) DummyEntity {
	return DummyEntity{
        Info: de.Info,
	}
}
