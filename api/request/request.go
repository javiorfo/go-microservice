package request

// Dummy represents the json object for creating a Dummy
type Dummy struct {
	Info string `json:"info" validate:"required"`
}
