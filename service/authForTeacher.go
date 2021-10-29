package service

import (
	todo "FactorialSchoolBook"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)
type tokenClaimsForTeacher struct{
	jwt.StandardClaims
	UserId int `json:"user_id"`
	Role string `json:"role"`
}

func (s *AuthService) CreateUserForTeacher(user todo.User) (int,error){
	user.Password=generatePasswordHash(user.Password)
	return s.repo.CreateUserForTeacher(user)
}
func (s *AuthService) GenerateTokenForTeacher(username, password string) (string,error){
	user,err:=s.repo.GetUserForTeacher(username,generatePasswordHash(password))
	if err!=nil{
		return "",err
	}
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,&tokenClaimsForTeacher{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt: time.Now().Unix(),
		},
		UserId: user.Id,
		Role: "Teacher",
	})
	return token.SignedString([]byte(signingKey))
}
func (s *AuthService) ParseTokenForTeacher(accessToken string) (todo.UserAuth,error){
	token,err := jwt.ParseWithClaims(accessToken,&tokenClaimsForTeacher{},func(token *jwt.Token)(interface{},error){
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil,errors.New("invalid signing method")
		}
		return []byte(signingKey),nil
	})
	if err!=nil{
		return todo.UserAuth{},err
	}
	claims,ok:= token.Claims.(*tokenClaimsForTeacher)
	if !ok{
		return todo.UserAuth{},errors.New("token claims are not of type *tokenClaims")
	}
	return todo.UserAuth{UserId: claims.UserId,Role: claims.Role},nil
}
