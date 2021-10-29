package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
func (h *Handler) getAllCoursesForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err=isTeacher(c)
	if err!=nil{
		return
	}
	courses,err := h.services.TodoCourse.GetAllForTeacher(id)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllCourseResponse{
		Data: courses,
	})
}
func (h *Handler) getCourseByIdForTeacher(c *gin.Context){
	userId,err := getUserId(c)
	if err!=nil{
		return
	}
	err=isTeacher(c)
	if err!=nil{
		return
	}
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	course,err := h.services.TodoCourse.GetByIdForTeacher(userId,id)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,course)
}

