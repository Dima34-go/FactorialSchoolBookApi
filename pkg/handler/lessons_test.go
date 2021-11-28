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

func TestHandler_getAllLessons(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoLesson, userId ,courseId int)
	testTable := []struct {
		name                string
		userInfo            todo.UserAuth
		courseId            int
		courseIdStr         string
		lessonId            int
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId int ) {
				s.EXPECT().GetAll(userId,courseId).Return([]todo.Lesson{{Id: 1,Title: "some title",
					StatusLesson: "проведено",AvailableForLearner: true, AvailableForTeacher: true}},nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `[{"id":1,"title":"some title","description":"","status_lesson":"проведено","available_for_learner":true,"available_for_teacher":true}]`,
		},
		{
			name:                "SERVICE ERROR",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 1,
			lessonId: 1,
			courseIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId int ) {
				s.EXPECT().GetAll(userId,courseId).Return(nil, errors.New("SERVICE ERROR"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"SERVICE ERROR"}`,
		},
		{
			name:                "INVALID CourseId FORMAT",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 0,
			lessonId: 1,
			courseIdStr: "rrr",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId int ) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid id param"}`,
		},
		{
			name:                "INVALID ROLE",
			userInfo:            todo.UserAuth{UserId: 1, Role: "wrong role"},
			courseId: 1,
			lessonId: 1,
			courseIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId int ) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user role is not learner"}`,
		},
		{
			name:                "INVALID UserID",
			userInfo:            todo.UserAuth{ Role: learnerRole},
			courseId: 1,
			lessonId: 1,
			courseIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId int ) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"user id is invalid type"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			lessons := mock_service.NewMockTodoLesson(c)
			testCase.mockBehavior(lessons, testCase.userInfo.UserId , testCase.courseId)
			services := &service.Service{TodoLesson: lessons}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/:id/allLessons",
				func(c *gin.Context) {
					if testCase.userInfo.UserId == 0 {
						c.Set(userCtx, "some wrong info")
					} else {
						c.Set(userCtx, testCase.userInfo.UserId)
					}

					c.Set(roleCtx, testCase.userInfo.Role)
				},
				handler.getAllLessons)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/"+testCase.courseIdStr+"/allLessons",
				nil)

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
func TestHandler_getLessonById(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTodoLesson, userId ,courseId ,lessonId int)
	testTable := []struct {
		name                string
		userInfo            todo.UserAuth
		courseId            int
		courseIdStr         string
		lessonId            int
		lessonIdStr         string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:                "OK",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			lessonIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {
				s.EXPECT().GetById(userId,courseId,lessonId).Return(todo.Lesson{Id: 1,Title: "some title",
					StatusLesson: "проведено",AvailableForLearner: true, AvailableForTeacher: true},nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: `{"id":1,"title":"some title","description":"","status_lesson":"проведено","available_for_learner":true,"available_for_teacher":true}`,
		},
		{
			name:                "SERVICE ERROR",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			lessonIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {
				s.EXPECT().GetById(userId,courseId,lessonId).Return(todo.Lesson{},errors.New("SERVICE PROBLEM"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"SERVICE PROBLEM"}`,
		},
		{
			name:                "INVALID CourseId PARAM",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 1,
			courseIdStr: "rrr",
			lessonId: 1,
			lessonIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid id param"}`,
		},
		{
			name:                "INVALID LessonId PARAM",
			userInfo:            todo.UserAuth{UserId: 1, Role: learnerRole},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			lessonIdStr: "rrr",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid lesson_id param"}`,
		},
		{
			name:                "INVALID UserId PARAM",
			userInfo:            todo.UserAuth{ Role: learnerRole},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			lessonIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user id is invalid type"}`,
		},
		{
			name:                "ROLE IS MISSING",
			userInfo:            todo.UserAuth{UserId: 1, Role: "role is missing"},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			lessonIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user role is not found"}`,
		},
		{
			name:                "UserID IS MISSING",
			userInfo:            todo.UserAuth{UserId: -1, Role: learnerRole},
			courseId: 1,
			courseIdStr: "1",
			lessonId: 1,
			lessonIdStr: "1",
			mockBehavior:        func(s *mock_service.MockTodoLesson, userId ,courseId,lessonId int ) {},
			expectedStatusCode:  401,
			expectedRequestBody: `{"message":"user id is not found"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init Deps
			c := gomock.NewController(t)
			defer c.Finish()

			lessons := mock_service.NewMockTodoLesson(c)
			testCase.mockBehavior(lessons, testCase.userInfo.UserId , testCase.courseId, testCase.lessonId)
			services := &service.Service{TodoLesson: lessons}
			handler := NewHandler(services)

			//Test Server
			gin.SetMode(gin.ReleaseMode)
			r := gin.New()
			r.GET("/:id/allLessons/:lesson_id",
				func(c *gin.Context) {
					if testCase.userInfo.UserId == 0 {
						c.Set(userCtx, "some wrong info")
					} else if  testCase.userInfo.UserId != -1{
						c.Set(userCtx, testCase.userInfo.UserId)
					}
					if !(testCase.userInfo.Role == "role is missing"){
						c.Set(roleCtx, testCase.userInfo.Role)
					}
				},
				handler.getLessonById)

			//Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/"+testCase.courseIdStr+"/allLessons/"+testCase.lessonIdStr,
				nil)

			//Perform Request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedRequestBody)
		})
	}
}
