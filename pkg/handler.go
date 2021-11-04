package handler

import (
	"FactorialSchoolBook/service"
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
func NewHandler(services *service.Service) *Handler{
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
				//позже будет убрано
				lessons.GET("/:lesson_id/get",h.getHomework)
			}
		}
	}
	//авторизация и регистрация для учителей
	authForTeacher:=router.Group(teacherRelative+"/auth")
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
				//позже будет убрано
				lessons.POST("/:lesson_id",h.addHomeworkForTeacher)
			}
		}
	}

	return router
}