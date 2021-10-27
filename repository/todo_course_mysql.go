package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CourseMysql struct {
	db *sqlx.DB
}
func NewCourseMysql(db *sqlx.DB) *CourseMysql{
	return &CourseMysql{db: db}
}
func (r *CourseMysql) GetAll(userId int) ([]todo.Course,error){
	var courses []todo.Course
	query:=fmt.Sprintf("SELECT курсы.Idкурса, название , описание  FROM %s INNER JOIN %s USING(IdКурса) WHERE IdУченика = ? ",courseTables,learnerCoursesTables)
	err:=r.db.Select(&courses,query,userId)
	return courses,err
}
func (r *CourseMysql) GetById(userId, courseId int) (todo.Course,error){
	var course todo.Course
	query:=fmt.Sprintf("SELECT Idкурса, название , описание  FROM %s INNER JOIN %s USING(IdКурса) WHERE IdУченика = ? AND IdКурса = ?",courseTables,learnerCoursesTables)
	err:=r.db.Get(&course,query,userId,courseId)
	return course,err
}