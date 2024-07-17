package dummy

import (
	"github.com/javiorfo/go-microservice/domain/model"
)

type Service struct{}

func (d *Service) FindById(id int) (*model.Dummy, error) {
	return &model.Dummy{Info: "info"}, nil
}
