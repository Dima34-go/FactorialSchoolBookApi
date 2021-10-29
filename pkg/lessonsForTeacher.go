package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func (h *Handler) getAllLessonsForTeacher(c *gin.Context){
	userId,err := getUserId(c)
	if err!=nil{
		return
	}
	err=isTeacher(c)
	if err!=nil{
		return
	}
	courseId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	lessons,err :=h.services.TodoLesson.GetAllForTeacher(userId,courseId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,lessons)
}
func (h *Handler) getLessonByIdForTeacher(c *gin.Context){
	userId,err := getUserId(c)
	if err!=nil{
		return
	}
	err=isTeacher(c)
	if err!=nil{
		return
	}
	courseId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	lessonId,err:=strconv.Atoi(c.Param("lesson_id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	lesson,err :=h.services.TodoLesson.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,lesson)
}