package lib

import (
	"github.com/gin-gonic/gin"
)

const defaultLimit = 100
const maxLimit = 1000

// Pagination contains pagination request
type Pagination struct {
	Limit int64 `form:"limit"`
	Page  int64 `form:"page" binding:"min=0"`
	Skip  int64 `json:"-"`
}

// Paginate validates pagination requests
func Paginate(c *gin.Context) (*Pagination, error) {
	p := new(Pagination)
	if err := c.ShouldBindQuery(p); err == nil {
		if p.Limit < 1 {
			p.Limit = defaultLimit
		}
		if p.Limit > 1000 {
			p.Limit = maxLimit
		}
		p.Skip = p.Limit * p.Page
	}
	return p, nil
}
