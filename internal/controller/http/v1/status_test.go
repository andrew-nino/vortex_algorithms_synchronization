package v1

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/service"
	service_mocks "github.com/andrew-nino/vtx_algorithms_synchronization/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_updateAlgorithmStatus(t *testing.T) {

	type mockBehavior func(r *service_mocks.MockAlgorithmStatus, client entity.AlgorithmStatus)

	tests := []struct {
		name                 string
		inputBody            string
		inputClient          entity.AlgorithmStatus
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBody: `{"client_id":32, "vwap": false, "twap": false, "hft": false}`,
			inputClient: entity.AlgorithmStatus{
				ClientID: 32,
				VWAP:     false,
				TWAP:     false,
				HFT:      false,
			},
			mockBehavior: func(r *service_mocks.MockAlgorithmStatus, newStatus entity.AlgorithmStatus) {
				r.EXPECT().UpdateStatus(newStatus).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"description":"status updated successfully","newStatus":{"client_id":32,"vwap":false,"twap":false,"hft":false}}`,
		},
		{
			name: "Wrong input",
			inputBody: `{"vwap": false, "twap": false, "hft": false}`,
			inputClient:          entity.AlgorithmStatus{},
			mockBehavior:         func(r *service_mocks.MockAlgorithmStatus, client entity.AlgorithmStatus) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"Key: 'AlgorithmStatus.ClientID' Error:Field validation for 'ClientID' failed on the 'required' tag"}`,
		},
		{
			name: "Service Error",
			inputBody: `{"client_id":32, "vwap": false, "twap": false, "hft": false}`,
			inputClient: entity.AlgorithmStatus{
				ClientID: 32,
				VWAP:     false,
				TWAP:     false,
				HFT:      false,
			},
			mockBehavior: func(r *service_mocks.MockAlgorithmStatus, client entity.AlgorithmStatus) {
				r.EXPECT().UpdateStatus(client).Return(errors.New("update status failed"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"update status failed"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockAlgorithmStatus(c)
			test.mockBehavior(repo, test.inputClient)

			services := &service.Service{AlgorithmStatus: repo}
			handler := Handler{services}

			r := gin.New()
			r.PUT("/update", handler.updateAlgorithmStatus)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/update",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
