package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
)

func (r *CourseMysql) GetAllForTeacher(userId int) ([]todo.Course,error){
	var courses []todo.Course
	query:=fmt.Sprintf("SELECT курсы.Idкурса, название , описание  FROM %s INNER JOIN %s USING(IdКурса) WHERE IdПреподавателя = ? ",courseTables,teacherCoursesTables)
	err:=r.db.Select(&courses,query,userId)
	return courses,err
}
func (r *CourseMysql) GetByIdForTeacher(userId, courseId int) (todo.Course,error){
	var course todo.Course
	query:=fmt.Sprintf("SELECT Idкурса, название , описание  FROM %s INNER JOIN %s USING(IdКурса) WHERE IdПреподавателя = ? AND IdКурса = ?",courseTables,teacherCoursesTables)
	err:=r.db.Get(&course,query,userId,courseId)
	return course,err
}