package handler

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/service"
	mock_service "FactorialSchoolBook/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user todo.User)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           todo.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: todo.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user todo.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name:      "Some data is missing",
			inputBody: `{"name":"Test","username":"test"}`,
			inputUser: todo.User{
				Name:     "Test",
				Username: "test",
			},
			mockBehavior:        func(s *mock_service.MockAuthorization, user todo.User) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Key: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}`,
		},
		{
			name:      "Problems on the service",
			inputBody: `{"name":"Test","username":"test","password":"qwerty"}`,
			inputUser: todo.User{
				Name:     "Test",
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user todo.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("wrong password"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"wrong password"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())

		})
	}
}
func TestHandler_signIn(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user todo.SignInInput)
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           todo.SignInInput
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"username":"test","password":"qwerty"}`,
			inputUser: todo.SignInInput{
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user todo.SignInInput) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("token", nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"token":"token"}`,
		},
		{
			name:      "wrong data",
			inputBody: `{"username":"test","password":"qwerty"}`,
			inputUser: todo.SignInInput{
				Username: "test",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user todo.SignInInput) {
				s.EXPECT().GenerateToken(user.Username, user.Password).Return("token", errors.New("wrong data"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"wrong data"}`,
		},
		{
			name:      "bad request",
			inputBody: `{"password":"qwerty"}`,
			inputUser: todo.SignInInput{
				Password: "qwerty",
			},
			mockBehavior:        func(s *mock_service.MockAuthorization, user todo.SignInInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"Key: 'SignInInput.Username' Error:Field validation for 'Username' failed on the 'required' tag"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.POST("/sign-in", handler.signIn)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/sign-in",
				bytes.NewBufferString(testCase.inputBody))

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)

		})
	}
}
