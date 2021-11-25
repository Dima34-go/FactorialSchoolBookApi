package handler

import (
	todo "FactorialSchoolBook"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUpForTeacher(c *gin.Context){
	var input todo.User
	if err:= c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	id,err:= h.services.Authorization.CreateUserForTeacher(input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signInForTeacher(c *gin.Context){
	var input signInInput
	if err:= c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	token,err:= h.services.Authorization.GenerateTokenForTeacher(input.Username,input.Password)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]interface{}{
		"token": token,
	})
}