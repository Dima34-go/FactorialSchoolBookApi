package handler

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/service"
	mock_service "FactorialSchoolBook/pkg/service/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_getAllCourses(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoCourse, courseId int)
	testTable := []struct {
		name                string
		userInfo            todo.UserAuth
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:     "OK",
			userInfo: todo.UserAuth{UserId: 1, Role: learnerRole},
			mockBehavior: func(s *mock_service.MockTodoCourse, userId int) {
				s.EXPECT().GetAll(userId).Return([]todo.Course{{Id: 1, Title: "some course", Description: "some description"}}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"data":[{"id":1,"title":"some course","description":"some description"}]}`,
		},
		{
			name:                "BAD ROLE",
			userInfo:            todo.UserAuth{UserId: 1, Role: "bad role"},
			mockBehavior:        func(s *mock_service.MockTodoCourse, userId int) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user role is not learner"}`,
		},
		{
			name:     "SERVER PROBLEMS",
			userInfo: todo.UserAuth{UserId: 1, Role: learnerRole},
			mockBehavior: func(s *mock_service.MockTodoCourse, userId int) {
				s.EXPECT().GetAll(userId).Return(nil, errors.New("some service problem"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"some service problem"}`,
		},
		{
			name:                "UserID BAD FORMAT",
			userInfo:            todo.UserAuth{Role: learnerRole},
			mockBehavior:        func(s *mock_service.MockTodoCourse, userId int) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user id is invalid type"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			courses := mock_service.NewMockTodoCourse(c)
			testCase.mockBehavior(courses, testCase.userInfo.UserId)

			services := &service.Service{TodoCourse: courses}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/allCourses",
				func(c *gin.Context) {
					if testCase.userInfo.UserId == 0 {
						c.Set(userCtx, "some wrong info")
					} else {
						c.Set(userCtx, testCase.userInfo.UserId)
					}

					c.Set(roleCtx, testCase.userInfo.Role)
				},
				handler.getAllCourses)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/allCourses",
				nil)

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
func TestHandler_getCourseId(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoCourse, userId int, courseId int)
	testTable := []struct {
		name                string
		userInfo            todo.UserAuth
		courseId            int
		courseIdStr         string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:        "OK",
			userInfo:    todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId:    1,
			courseIdStr: "1",
			mockBehavior: func(s *mock_service.MockTodoCourse, userId int, courseId int) {
				s.EXPECT().GetById(userId, courseId).Return(todo.Course{Id: 1, Title: "some course", Description: "some description"}, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"title":"some course","description":"some description"}`,
		},
		{
			name:                "BAD ROLE",
			userInfo:            todo.UserAuth{UserId: 1, Role: "bad role"},
			courseId:            1,
			courseIdStr:         "1",
			mockBehavior:        func(s *mock_service.MockTodoCourse, userId int, courseId int) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user role is not learner"}`,
		},
		{
			name:        "SERVICE ERROR",
			userInfo:    todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId:    1,
			courseIdStr: "1",
			mockBehavior: func(s *mock_service.MockTodoCourse, userId int, courseId int) {
				s.EXPECT().GetById(userId, courseId).Return(todo.Course{}, errors.New("some service problem"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"some service problem"}`,
		},
		{
			name:                "UserID BAD FORMAT",
			userInfo:            todo.UserAuth{Role: learnerRole},
			courseId:            1,
			courseIdStr:         "1",
			mockBehavior:        func(s *mock_service.MockTodoCourse, userId int, courseId int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"user id is invalid type"}`,
		},
		{
			name:                "CourseID BAD FORMAT",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseIdStr:         "courseId",
			mockBehavior:        func(s *mock_service.MockTodoCourse, userId int, courseId int) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid id param"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			courses := mock_service.NewMockTodoCourse(c)
			testCase.mockBehavior(courses, testCase.userInfo.UserId, testCase.courseId)

			services := &service.Service{TodoCourse: courses}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/course/:id",
				func(c *gin.Context) {
					if testCase.userInfo.UserId == 0 {
						c.Set(userCtx, "some wrong info")
					} else {
						c.Set(userCtx, testCase.userInfo.UserId)
					}

					c.Set(roleCtx, testCase.userInfo.Role)
				},
				handler.getCourseById)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/course/"+testCase.courseIdStr,
				nil)

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
