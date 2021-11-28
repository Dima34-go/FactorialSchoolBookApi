package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) getHomeTask(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	err = isLearner(c)
	if err != nil {
		return
	}
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lessonId, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lesson_id param")
		return
	}
	homeTask, err := h.services.TodoLessonInfo.GetHomeTask(id, courseId, lessonId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, homeTask)
}
func (h *Handler) getHomework(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	err = isLearner(c)
	if err != nil {
		return
	}
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lessonId, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lesson_id param")
		return
	}
	homework, err := h.services.TodoLessonInfo.GetHomework(id, courseId, lessonId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, homework)
}
func (h *Handler) postHomework(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	err = isLearner(c)
	if err != nil {
		return
	}
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lessonId, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lesson_id param")
		return
	}
	err = h.services.TodoLessonInfo.PostHomework(id, courseId, lessonId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
func (h *Handler) getAttendance(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	err = isLearner(c)
	if err != nil {
		return
	}
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lessonId, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lesson_id param")
		return
	}
	attendance, err := h.services.TodoLessonInfo.GetAttendance(id, courseId, lessonId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, attendance)
}
func (h *Handler) postAttendance(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}
	err = isLearner(c)
	if err != nil {
		return
	}
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	lessonId, err := strconv.Atoi(c.Param("lesson_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid lesson_id param")
		return
	}
	err = h.services.TodoLessonInfo.PostAttendance(id, courseId, lessonId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}
