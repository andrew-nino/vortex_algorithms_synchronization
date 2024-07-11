package v1

import (
	"net/http"
	"strconv"

	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
}

// Adding a new client. If successful, we get the client id
func (h *Handler) addClient(c *gin.Context) {
	newClient := entity.Client{}

	if err := c.BindJSON(&newClient); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clientID, err := h.services.AddClient(newClient)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "client add failed")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: clientID})
}

// Update client data. If successful, we get the client id
func (h *Handler) updateClient(c *gin.Context) {
	updateClient := entity.Client{}

	if err := c.BindJSON(&updateClient); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	clientID, err := h.services.UpdateClient(updateClient)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "client update failed")
		return
	}
	c.JSON(http.StatusOK, response{Message: "success", ID: clientID})

}

// Deleting client data.
func (h *Handler) deleteClient(c *gin.Context) {

	paramStr := c.Query("client_id")
	if paramStr == "" {
		newErrorResponse(c, http.StatusBadRequest, "client_id is required")
		return
	}
	clientID, err := strconv.Atoi(paramStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "client_id must be an integer")
		return
	}
	err = h.services.DeleteClient(clientID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "client delete failed")
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "success"})

}
