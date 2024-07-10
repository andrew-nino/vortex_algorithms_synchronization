package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "managerID"
)

// User validation to determine access level.
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		log.Errorf("AuthMiddleware.userIdentity:  %v", ErrEmptyAuthHeader)
		newErrorResponse(c, http.StatusUnauthorized, ErrEmptyAuthHeader.Error())
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Errorf("AuthMiddleware.userIdentity: bearerToken: %v", ErrInvalidAuthHeader)
		newErrorResponse(c, http.StatusUnauthorized, ErrInvalidAuthHeader.Error())
		return
	}

	if len(headerParts[1]) == 0 {
		log.Errorf("AuthMiddleware.userIdentity: emtyToken: %v", ErrEmptyAythToken)
		newErrorResponse(c, http.StatusUnauthorized, ErrEmptyAythToken.Error())
		return
	}

	managerID, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		log.Errorf("AuthMiddleware.userIdentity: CannotParseToken: %v", ErrCannotParseToken)
		newErrorResponse(c, http.StatusUnauthorized, ErrCannotParseToken.Error())
		return
	}

	c.Set(userCtx, managerID)
}

func getManagerId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
