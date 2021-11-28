package handler

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/service"
	mock_service "FactorialSchoolBook/pkg/service/mocks"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_userIdentity(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, token string)
	testTable := []struct {
		name                 string
		headerName           string
		headerValue          string
		token                string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			headerName:  "Authorization",
			headerValue: "Bearer token",
			token:       "token",
			mockBehavior: func(s *mock_service.MockAuthorization, token string) {
				s.EXPECT().ParseToken(token).Return(todo.UserAuth{UserId: 1, Role: learnerRole}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: "1 " + learnerRole,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()
			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.token)
			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			//Test Server
			r := gin.New()
			r.GET("/protected", handler.userIdentity, func(c *gin.Context) {
				id, _ := c.Get(userCtx)
				role, _ := c.Get(roleCtx)
				c.String(200, fmt.Sprintf("%d %s", id.(int), role))
			})

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/protected", nil)
			req.Header.Set(testCase.headerName, testCase.headerValue)

			//Make Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)

		})
	}
}
