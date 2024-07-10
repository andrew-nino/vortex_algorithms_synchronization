package v1

import (
	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {

	var input entity.Manager

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateManager(input)
	if err != nil {
		log.Debugf("error when registering manager : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "internal server error")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	ManagerName string `json:"managername" binding:"required" example:"Manager"`
	Password    string `json:"password" binding:"required" example:"qwerty"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.SignIn(input.ManagerName, input.Password)
	if err != nil {
		log.Debugf("error during mamager verification : %s", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
