package handler

import (
	"FactorialSchoolBook/service"
	"github.com/gin-gonic/gin"
)

type Handler struct{
	services *service.Service
}
func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router:=gin.New()
	auth:=router.Group("/auth")
	{
		auth.POST("/sign-up",h.signUp)
		auth.POST("/sign-in",h.signIn)
	}
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
			}
		}
	}
	return router
}