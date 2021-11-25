package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthMysql struct{
	db *sqlx.DB
}
func NewAuthMysql(db *sqlx.DB) *AuthMysql {
	return &AuthMysql{db: db}
}
func (r *AuthMysql) CreateUser(user todo.User) (int,error){
	var id int
	tx,err :=r.db.Begin()
	query:= fmt.Sprintf("INSERT INTO %s (имя, логин, пароль, отчество, фамилия) values (?, ?, ?, ?, ?)", learnerTables)
	_,err =tx.Query(query, user.Name, user.Username, user.Password, user.Patronymic, user.Surname)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}
	query = fmt.Sprintf("SELECT idученика FROM %s ORDER BY idученика DESC LIMIT 1;", learnerTables)
	row:=tx.QueryRow(query)
	if err=row.Scan(&id);err!=nil{
		tx.Rollback()
		return 0,err
	}
	return id,tx.Commit()
}
func (r *AuthMysql) GetUser(username, password string) (todo.User,error){
	var user todo.User
	query:=fmt.Sprintf("SELECT idученика from %s where пароль = ? AND логин = ?  ", learnerTables)
	err:=r.db.Get(&user,query,password,username)
	return user,err
}