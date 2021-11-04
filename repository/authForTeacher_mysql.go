package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
)

func (r *AuthMysql) CreateUserForTeacher(user todo.User) (int,error){
	var id int
	tx,err:=r.db.Begin()
	query:= fmt.Sprintf("INSERT INTO %s (имя, логин, пароль, отчество, фамилия) values (?, ?, ?, ?, ?)", teachersTables)
	_,err =tx.Query(query, user.Name, user.Username, user.Password, user.Patronymic, user.Surname)
	if err!=nil{
		tx.Rollback()
		return 0,err
	}
	query = fmt.Sprintf("SELECT idпреподавателя FROM %s ORDER BY idпреподавателя DESC LIMIT 1;",teachersTables)
	row:=tx.QueryRow(query)
	if err=row.Scan(&id);err!=nil{
		tx.Rollback()
		return 0,err
	}
	return id,tx.Commit()
}
func (r *AuthMysql) GetUserForTeacher(username, password string) (todo.Teacher,error){
	var user todo.Teacher
	query:=fmt.Sprintf("SELECT idПреподавателя from %s where пароль = ? AND логин = ?  ",teachersTables)
	err:=r.db.Get(&user,query,password,username)
	return user,err
}