package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct{
	db *sqlx.DB
}
func NewAuthMysql(db *sqlx.DB) *AuthMysql{
	return &AuthMysql{db: db}
}
func (r *AuthMysql) CreateUser(user todo.User) (int,error){
	var id int
	query:= fmt.Sprintf("INSERT INTO %s (имя, логин, пароль, отчество, фамилия) values (?, ?, ?, ?, ?)", learnerTables)
	_,err:=r.db.Query(query, user.Name, user.Username, user.Password, user.Patronymic, user.Surname)
	if err!=nil{
		return 0,err
	}
	query = fmt.Sprintf("SELECT idученика FROM %s ORDER BY idученика DESC LIMIT 1;",learnerTables)
	row:=r.db.QueryRow(query)
	if err=row.Scan(&id);err!=nil{
		return 0,err
	}
	return id,nil
}
func (r *AuthMysql) GetUser(username, password string) (todo.User,error){
	var user todo.User
	query:=fmt.Sprintf("SELECT idученика from %s where пароль = ? AND логин = ?  ",learnerTables)
	err:=r.db.Get(&user,query,password,username)
	return user,err
}