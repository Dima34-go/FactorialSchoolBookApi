package handler

import (
	todo "FactorialSchoolBook"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createCourse(c *gin.Context){

}
func (h *Handler) deleteCourse(c *gin.Context){

}
func (h *Handler) updateCourse(c *gin.Context){

}
type getAllCourseResponse struct{
	Data []todo.Course `json:"data"`
}
func (h *Handler) getAllCourses(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	courses,err := h.services.TodoCourse.GetAll(id)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
	}
	c.JSON(http.StatusOK,getAllCourseResponse{
		Data: courses,
	})
}
func (h *Handler) getCourseById(c *gin.Context){
	userId,err := getUserId(c)
	if err!=nil{
		return
	}
	id,err:=strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	course,err := h.services.TodoCourse.GetById(userId,id)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
	}
	c.JSON(http.StatusOK,course)
}

