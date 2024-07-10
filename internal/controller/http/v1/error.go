package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidAuthHeader = fmt.Errorf("invalid auth header")
	ErrEmptyAuthHeader   = fmt.Errorf("empty auth header")
	ErrEmptyAythToken    = fmt.Errorf("token is empty")
	ErrCannotParseToken  = fmt.Errorf("cannot parse token")
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
