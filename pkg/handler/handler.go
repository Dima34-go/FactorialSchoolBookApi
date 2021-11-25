package handler

import (
	"FactorialSchoolBook/pkg/service"
	"github.com/gin-gonic/gin"
)
const(
	teacherRelative= "/teacher"
	teacherRole="Teacher"
	learnerRole="Learner"
)
type Handler struct{
	services *service.Service
}
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router:=gin.New()
	//авторизация и регистрация для учеников
	auth:=router.Group("/auth")
	{
		auth.POST("/sign-up",h.signUp)
		auth.POST("/sign-in",h.signIn)
	}
	//api для ученика
	api:=router.Group("/api",h.userIdentity)
	{
		courses:=api.Group("/courses")
		{
			courses.GET("/",h.getAllCourses)
			courses.GET("/:id",h.getCourseById)
			lessons:=courses.Group(":id/lessons")
			{
				lessons.GET("/",h.getAllLessons)
				lessons.GET("/:lesson_id",h.getLessonById)
				lessonInfo:=lessons.Group(":lesson_id")
				{
					lessonInfo.GET("/homeTask",h.getHomeTask)
					lessonInfo.GET("/homework",h.getHomework)
					lessonInfo.POST("/homework",h.postHomework)
					lessonInfo.GET("/attendance",h.getAttendance)
					lessonInfo.POST("/attendance",h.postAttendance)
				}
			}
		}
	}
	//авторизация и регистрация для учителей
	authForTeacher:=router.Group(teacherRelative +"/auth")
	{
		authForTeacher.POST("/sign-up",h.signUpForTeacher)
		authForTeacher.POST("/sign-in",h.signInForTeacher)
	}
	//api для учителей
	apiForTeacher:=router.Group(teacherRelative+"/api",h.userIdentityForTeacher)
	{
		courses:=apiForTeacher.Group("/courses")
		{
			courses.GET("/",h.getAllCoursesForTeacher)
			courses.GET("/:id",h.getCourseByIdForTeacher)
			lessons:=courses.Group(":id/lessons")
			{
				lessons.GET("/",h.getAllLessonsForTeacher)
				lessons.POST("/", h.createLessonForTeacher)
				lessons.GET("/:lesson_id",h.getLessonByIdForTeacher)
				lessonInfo:=lessons.Group(":lesson_id")
				{
					lessonInfo.GET("/homeTask",h.getHomeTaskForTeacher)
					lessonInfo.POST("/homeTask",h.postHomeTaskForTeacher)
					lessonInfo.GET("/attendance",h.getAttendanceForTeacher)
					lessonInfo.POST("/attendance",h.postAttendanceForTeacher)
					lessonInfo.POST("/start",h.startLessonForTeacher)
					lessonInfo.POST("/finish",h.finishLessonForTeacher)
					lessonInfo.POST("/checkHomework",h.checkHomeworkForTeacher)
					lessonInfo.POST("/checkAttendance",h.checkAttendanceForTeacher)
				}
			}
		}
	}

	return router
}