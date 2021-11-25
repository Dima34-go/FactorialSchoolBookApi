package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createLesson(c *gin.Context){

}
func (h *Handler) deleteLesson(c *gin.Context){

}
func (h *Handler) updateLesson(c *gin.Context){

}
func (h *Handler) getAllLessons(c *gin.Context){
	userId,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isLearner(c)
	if err!=nil{
		return
	}
	courseId,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	lessons,err :=h.services.TodoLesson.GetAll(userId,courseId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,lessons)
}
func (h *Handler) getLessonById(c *gin.Context){
	userId,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isLearner(c)
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
	lesson,err :=h.services.TodoLesson.GetById(userId,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,lesson)
}