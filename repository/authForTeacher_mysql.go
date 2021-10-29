package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
)

func (r *AuthMysql) CreateUserForTeacher(user todo.User) (int,error){
	var id int
	query:= fmt.Sprintf("INSERT INTO %s (имя, логин, пароль, отчество, фамилия) values (?, ?, ?, ?, ?)", teachersTables)
	_,err:=r.db.Query(query, user.Name, user.Username, user.Password, user.Patronymic, user.Surname)
	if err!=nil{
		return 0,err
	}
	query = fmt.Sprintf("SELECT idпреподавателя FROM %s ORDER BY idпреподавателя DESC LIMIT 1;",teachersTables)
	row:=r.db.QueryRow(query)
	if err=row.Scan(&id);err!=nil{
		return 0,err
	}
	return id,nil
}
func (r *AuthMysql) GetUserForTeacher(username, password string) (todo.Teacher,error){
	var user todo.Teacher
	query:=fmt.Sprintf("SELECT idПреподавателя from %s where пароль = ? AND логин = ?  ",teachersTables)
	err:=r.db.Get(&user,query,password,username)
	return user,err
}