package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) userIdentityForTeacher(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	//parse token
	user, err := h.services.Authorization.ParseTokenForTeacher(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	c.Set(userCtx, user.UserId)
	c.Set(roleCtx, user.Role)
}
func isTeacher(c *gin.Context) error {
	userRole, err := getUserRole(c)
	if err != nil {
		return err
	}
	if userRole != teacherRole {
		newErrorResponse(c, http.StatusInternalServerError, "user role is not teacher")
		return errors.New("user role is not " + teacherRole)
	}
	return nil
}
