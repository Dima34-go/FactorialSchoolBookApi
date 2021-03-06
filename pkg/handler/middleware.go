package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "role"
)

func (h *Handler) userIdentity(c *gin.Context) {
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
	user, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}
	c.Set(userCtx, user.UserId)
	c.Set(roleCtx, user.Role)
}
func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id is not found")
		return 0, errors.New("user id is not found")
	}
	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user id is invalid type")
		return 0, errors.New("user id is invalid type")
	}
	return idInt, nil
}
func getUserRole(c *gin.Context) (string, error) {
	role, ok := c.Get(roleCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user role is not found")
		return "", errors.New("user id is not found")
	}
	roleString, ok := role.(string)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "user role is invalid type")
		return "", errors.New("user id is invalid type")
	}
	return roleString, nil
}
func isLearner(c *gin.Context) error {
	userRole, err := getUserRole(c)
	if err != nil {
		return err
	}
	if userRole != learnerRole {
		newErrorResponse(c, http.StatusUnauthorized, "user role is not learner")
		return errors.New("user role is not " + learnerRole)
	}
	return nil
}
