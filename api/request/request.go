package request

type Dummy struct {
	Info string `json:"info" validate:"required"`
}
