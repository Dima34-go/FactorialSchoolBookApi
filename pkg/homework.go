package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
type File struct{
	Name string `json:"name"`
}
func (h *Handler) getHomework(c *gin.Context){
	_,err := getUserId(c)
	if err!=nil{
		return
	}
	err=isLearner(c)
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
	//получение имени файла
	var FileForGet File
	if err=c.BindJSON(&FileForGet);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
	}
	//получение файла из директории и его отправка по http
	c.Header("Content-Type","multipart/form-data")
	c.File("static/"+c.Param("id")+"/"+c.Param("lesson_id")+"/"+FileForGet.Name)
}