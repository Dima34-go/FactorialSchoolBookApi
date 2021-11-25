package handler

import (
	todo "FactorialSchoolBook"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getHomeTaskForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	homeTask,err:=h.services.TodoLessonInfo.GetHomeTaskForTeacher(id,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,homeTask)
}

func (h *Handler) postHomeTaskForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	err =h.services.TodoLessonInfo.PostHomeTaskForTeacher(id,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]string{
		"status": "ok",
	})
}

func (h *Handler) getAttendanceForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	allInfo,err :=h.services.TodoLessonInfo.GetAttendanceForTeacher(id,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,allInfo)
}
func (h *Handler) postAttendanceForTeacher(c *gin.Context){

}
func (h *Handler) startLessonForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	err =h.services.TodoLessonInfo.StartLessonForTeacher(id,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]string{
		"status": "ok",
	})
}
func (h *Handler) finishLessonForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	err =h.services.TodoLessonInfo.FinishLessonForTeacher(id,courseId,lessonId)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]string{
		"status": "ok",
	})
}
func (h *Handler) checkHomeworkForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	var input []todo.LearnerHomework
	if err:= c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	err = h.services.TodoLessonInfo.CheckHomeworkForTeacher(id,courseId,lessonId,input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]string{
		"status": "ok",
	})
}
func (h *Handler) checkAttendanceForTeacher(c *gin.Context){
	id,err := getUserId(c)
	if err!=nil{
		return
	}
	err= isTeacher(c)
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
		newErrorResponse(c,http.StatusBadRequest,"invalid lesson_id param")
		return
	}
	var input []todo.LearnerAttendance
	if err:= c.BindJSON(&input);err!=nil{
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}
	err = h.services.TodoLessonInfo.CheckAttendanceForTeacher(id,courseId,lessonId,input)
	if err!=nil{
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return
	}
	c.JSON(http.StatusOK,map[string]string{
		"status": "ok",
	})
}

