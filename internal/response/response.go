package response

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/javiorfo/go-microservice/internal/response/codes"
	"github.com/javiorfo/go-microservice/internal/tracing"
)

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
}

type RestResponsePagination[T any] struct {
	Pagination PaginationResponse `json:"pagination"`
	Elements   []T                `json:"elements"`
}

type restResponseError struct {
	Errors []ResponseError `json:"errors"`
}

func (rre *restResponseError) AddError(c *fiber.Ctx, re ResponseError) *restResponseError {
    log.Errorf("%s Code: %s Message: %s", tracing.LogTraceAndSpan(c), re.Code, re.Message)
	rre.Errors = append(rre.Errors, re)
	return rre
}

func (rre *restResponseError) AddErrorWithCodeAndMsg(c *fiber.Ctx, code, msg string) *restResponseError {
    return rre.AddError(c, ResponseError{code, msg})
}

func NewRestResponseError(c *fiber.Ctx, re ResponseError) *restResponseError {
    log.Errorf("%s Code: %s Message: %s", tracing.LogTraceAndSpan(c), re.Code, re.Message)
	return &restResponseError{
		Errors: []ResponseError{re},
	}
}

func NewRestResponseErrorWithCodeAndMsg(c *fiber.Ctx, code, msg string) *restResponseError {
    return NewRestResponseError(c, ResponseError{code, msg})
}

func InternalServerError(c *fiber.Ctx, msg string) *restResponseError {
    return NewRestResponseError(c, ResponseError{codes.INTERNAL_ERROR, msg})
}
