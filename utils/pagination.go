package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Page   int
	Limit  int
	Offset int
}

func NewPaginationFromQuery(ctx *gin.Context) (*Pagination, error) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		return nil, errors.New("invalid page number")
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 50 {
		return nil, errors.New("invalid limit number")
	}

	return &Pagination{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}, nil
}
