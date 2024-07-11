package v1

import (
	"net/http"

	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/gin-gonic/gin"
)

// Update algorithm statuses for clients. If the update is successful, we receive a confirmation with new statuses for the client.
func (h *Handler) updateAlgorithmStatus(c *gin.Context) {

	newStatus := entity.AlgorithmStatus{}

	if err := c.BindJSON(&newStatus); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.UpdateStatus(newStatus)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "update status failed")
		return
	}

	type fullResponse struct {
		Description string                 `json:"description"`
		NewStatus   entity.AlgorithmStatus `json:"newStatus"`
	}
	fullresponse := fullResponse{
		Description: "status updated successfully",
		NewStatus:   newStatus,
	}
	c.JSON(http.StatusOK, fullresponse)
}
