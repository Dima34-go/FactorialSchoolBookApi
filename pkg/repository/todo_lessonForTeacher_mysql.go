package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
)

func (r *LessonMysql) GetAllForTeacher(userId, courseId int) ([]todo.Lesson,error){
	var lessons []todo.Lesson
	query:=fmt.Sprintf(`SELECT Idзанятия, занятия.название, занятия.описание , СтатусЗанятия , ДоступностьДляУченика , ДоступностьДляПреподавателя
                              FROM  %s INNER JOIN %s USING(Idкурса)
                              INNER JOIN %s USING(Idкурса) WHERE idКурса = ? AND idПреподавателя = ?`, lessonTables, courseTables, teacherCoursesTables)
	if err:=r.db.Select(&lessons,query,courseId,userId);err!=nil{
		return nil,err
	}
	return lessons,nil
}
func (r *LessonMysql) GetByIdForTeacher(userId, courseId ,lessonId int) (todo.Lesson,error){
	var lesson todo.Lesson
	query:=fmt.Sprintf(`SELECT Idзанятия, занятия.название, занятия.описание , СтатусЗанятия , ДоступностьДляУченика , ДоступностьДляПреподавателя
                              FROM  %s INNER JOIN %s USING(Idкурса)
                              INNER JOIN %s USING(Idкурса) 
                              WHERE idКурса = ? AND idПреподавателя = ? AND IdЗанятия = ? AND ДоступностьДляПреподавателя = 1`, lessonTables, courseTables, teacherCoursesTables)
	if err:=r.db.Get(&lesson,query,courseId,userId,lessonId);err!=nil{
		return todo.Lesson{},err
	}
	return lesson,nil
}
func (r *LessonMysql) CreateLessonForTeacher(lesson todo.Lesson,courseId  int) (int,error){
	var id int
	tx,err:=r.db.Begin()
	query:= fmt.Sprintf("INSERT INTO %s (idКурса ,название , описание , статусЗанятия) values ( ?, ?, ?, ?)", lessonTables)
	_,err =tx.Query(query,courseId,lesson.Title,lesson.Description, "Закрыто")
	if err!=nil{
		tx.Rollback()
		return 0,err
	}
	query = fmt.Sprintf("SELECT idзанятия FROM %s ORDER BY idзанятия DESC LIMIT 1;", lessonTables)
	row:=tx.QueryRow(query)
	if err=row.Scan(&id);err!=nil{
		tx.Rollback()
		return 0,err
	}
	return id,tx.Commit()
}


