package in

type FindByIdUseCase[T any] interface {
    FindById(id int) (T, error)
}
