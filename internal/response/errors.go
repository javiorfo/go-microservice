package response

type restResponseError struct {
	Errors []ResponseError `json:"errors"`
}

func NewRestResponseError(re ResponseError) *restResponseError {
	return &restResponseError{
		Errors: []ResponseError{re},
	}
}

func NewRestResponseErrorWithCodeAndMsg(code, msg string) *restResponseError {
	return &restResponseError{
		Errors: []ResponseError{{code, msg}},
	}
}

func (rre *restResponseError) AddError(re ResponseError) *restResponseError {
	rre.Errors = append(rre.Errors, re)
	return rre
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	PageNumber uint `json:"pageNumber"`
	PageSize   uint `json:"pageSize"`
	Total      uint `json:"total"`
}

type RestResponsePagination[T any] struct {
	Pagination PaginationResponse `json:"pagination"`
	Elements   []T                `json:"elements"`
}
