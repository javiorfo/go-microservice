package pagination

import (
	"errors"
	"strconv"

	"github.com/javiorfo/go-microservice/internal/response"
)

type Page struct {
	Page      int
	Size      int
	SortBy    string
	SortOrder string
}

func ValidateAndGetPage(page, size, sortBy, sortOrder string) (*Page, error) {
	p := Page{SortBy: sortBy}

	if pageInt, err := strconv.Atoi(page); err != nil {
		return nil, errors.New("'page' parameter must be a number")
	} else {
		p.Page = pageInt
	}

	if sizeInt, err := strconv.Atoi(size); err != nil {
		return nil, errors.New("'size' parameter must be a number")
	} else {
		p.Size = sizeInt
	}

	if sortOrder == "asc" || sortOrder == "desc" {
		p.SortOrder = sortOrder
		return &p, nil
	} else {
		return nil, errors.New("'sortOrder' parameter must be 'asc' or 'desc'")
	}
}

func Paginator(p Page, total int) response.PaginationResponse {
	return response.PaginationResponse{
		PageNumber: p.Page,
		PageSize:   p.Size,
		Total:      total,
	}
}

/* public class Paginator {
    public record Pair<T>(PaginationResponse pagination, List<T> results) {
        public Pair {
            pagination = Objects.requireNonNull(pagination);
        }
    }

    public static <T, R> Pair<R> create(Page<T> page, Function<T, R> mapper) {
        var content = page.getContent().stream().map(mapper).toList();
        var pageNumber = page.getPageable().getPageNumber();
        var pageSize = page.getPageable().getPageSize();
        var total = page.getTotalElements();
        var pagination = new PaginationResponse(pageNumber, pageSize, total);
        return new Pair<R>(pagination, content);
    }
} */
