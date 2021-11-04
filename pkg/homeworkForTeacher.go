package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) addHomeworkForTeacher(c *gin.Context){
	_,err := getUserId(c)
	if err!=nil{
		return
	}
	err=isTeacher(c)
	if err!=nil{
		return
	}
	_ ,err =strconv.Atoi(c.Param("id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}
	_ ,err =strconv.Atoi(c.Param("lesson_id"))
	if err!=nil{
		newErrorResponse(c,http.StatusBadRequest,"invalid id param")
		return
	}

	//new code
	mR, err := c.Request.MultipartReader()
	if err != nil {
		newErrorResponse(c,http.StatusBadRequest,"request not consist files")
		return
	}
	err=h.services.TodoHomework.AddHomeworkForTeacher(c.Param("id"),c.Param("lesson_id"),mR)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	//add new info on db
	c.JSON(http.StatusOK,map[string]interface{}{
		"status": "ok",
	})
}