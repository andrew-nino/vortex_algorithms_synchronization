package v1

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/andrew-nino/vtx_algorithms_synchronization/entity"
	"github.com/andrew-nino/vtx_algorithms_synchronization/internal/service"
	service_mocks "github.com/andrew-nino/vtx_algorithms_synchronization/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestHandler_addClient(t *testing.T) {

	type mockBehavior func(r *service_mocks.MockClient, client entity.Client)

	tests := []struct {
		name                 string
		inputBody            string
		inputClient          entity.Client
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBody: `{"client_id":126, "client_name":"Bulba", "version":1, "image":"test", "cpu":"intel",
						"memory":"32", "priority":2.5, "need_restart": true, "spawned_at":"2024-07-01"}`,
			inputClient: entity.Client{
				ID:          126,
				ClientName:  "Bulba",
				Version:     1,
				Image:       "test",
				CPU:         "intel",
				Memory:      "32",
				Priority:    2.5,
				NeedRestart: true,
				SpawnedAt:   "2024-07-01",
			},
			mockBehavior: func(r *service_mocks.MockClient, client entity.Client) {
				r.EXPECT().AddClient(client).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"success","id":1}`,
		},
		{
			name: "Wrong Input",
			inputBody: `{"client_id": 126, "client_name": "Bulba", "version": 1,  "image": "test","cpu": "intel",
						 "memory": "32",  "priority": 2.5, "need_restart": true}`,
			inputClient:          entity.Client{},
			mockBehavior:         func(r *service_mocks.MockClient, client entity.Client) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Service Error",
			inputBody: `{"client_id":126, "client_name":"Bulba", "version":1, "image":"test", "cpu":"intel",
						"memory":"32", "priority":2.5, "need_restart": true, "spawned_at":"2024-07-01"}`,
			inputClient: entity.Client{
				ID:          126,
				ClientName:  "Bulba",
				Version:     1,
				Image:       "test",
				CPU:         "intel",
				Memory:      "32",
				Priority:    2.5,
				NeedRestart: true,
				SpawnedAt:   "2024-07-01",
			},
			mockBehavior: func(r *service_mocks.MockClient, client entity.Client) {
				r.EXPECT().AddClient(client).Return(0, errors.New("client add failed"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"client add failed"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockClient(c)
			test.mockBehavior(repo, test.inputClient)

			services := &service.Service{Client: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/client", handler.addClient)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/client",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_updateClient(t *testing.T) {

	type mockBehavior func(r *service_mocks.MockClient, client entity.Client)

	tests := []struct {
		name                 string
		inputBody            string
		inputClient          entity.Client
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			inputBody: `{"client_id":126, "client_name":"Bulba", "version":1, "image":"test", "cpu":"intel",
						"memory":"32", "priority":2.5, "need_restart": true, "spawned_at":"2024-07-01"}`,
			inputClient: entity.Client{
				ID:          126,
				ClientName:  "Bulba",
				Version:     1,
				Image:       "test",
				CPU:         "intel",
				Memory:      "32",
				Priority:    2.5,
				NeedRestart: true,
				SpawnedAt:   "2024-07-01",
			},
			mockBehavior: func(r *service_mocks.MockClient, client entity.Client) {
				r.EXPECT().UpdateClient(client).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"success","id":1}`,
		},
		{
			name: "Wrong Input",
			inputBody: `{"client_id": 126, "client_name": "Bulba", "version": 1,  "image": "test","cpu": "intel",
						 "memory": "32",  "priority": 2.5, "need_restart": true}`,
			inputClient:          entity.Client{},
			mockBehavior:         func(r *service_mocks.MockClient, client entity.Client) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Service Error",
			inputBody: `{"client_id":126, "client_name":"Bulba", "version":1, "image":"test", "cpu":"intel",
						"memory":"32", "priority":2.5, "need_restart": true, "spawned_at":"2024-07-01"}`,
			inputClient: entity.Client{
				ID:          126,
				ClientName:  "Bulba",
				Version:     1,
				Image:       "test",
				CPU:         "intel",
				Memory:      "32",
				Priority:    2.5,
				NeedRestart: true,
				SpawnedAt:   "2024-07-01",
			},
			mockBehavior: func(r *service_mocks.MockClient, client entity.Client) {
				r.EXPECT().UpdateClient(client).Return(0, errors.New("client update failed"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"client update failed"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockClient(c)
			test.mockBehavior(repo, test.inputClient)

			services := &service.Service{Client: repo}
			handler := Handler{services}

			r := gin.New()
			r.POST("/update", handler.updateClient)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/update",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deleteClient(t *testing.T) {

	type mockBehavior func(r *service_mocks.MockClient, clientID int)

	tests := []struct {
		name                 string
		inputQuery           string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "Ok",
			inputQuery: "32",
			mockBehavior: func(r *service_mocks.MockClient, clientID int) {
				r.EXPECT().DeleteClient(clientID).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"message":"success"}`,
		},
		{
			name:                 "Empty query",
			inputQuery:           "",
			mockBehavior:         func(r *service_mocks.MockClient, clientID int) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"client_id is required"}`,
		},
		{
			name:                 "Query is not an integer",
			inputQuery:           "Oops",
			mockBehavior:         func(r *service_mocks.MockClient, clientID int) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"client_id must be an integer"}`,
		},
		{
			name:       "Service Error",
			inputQuery: "32",
			mockBehavior: func(r *service_mocks.MockClient, clientID int) {
				r.EXPECT().DeleteClient(clientID).Return(errors.New("client delete failed"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"client delete failed"}`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := service_mocks.NewMockClient(c)

			if test.inputQuery != "" {
				paramInt, _ := strconv.Atoi(test.inputQuery)
				test.mockBehavior(repo, paramInt)
			}

			services := &service.Service{Client: repo}
			handler := Handler{services}

			r := gin.New()
			r.DELETE("/delete", handler.deleteClient)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/delete?client_id="+test.inputQuery, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})

	}
}
