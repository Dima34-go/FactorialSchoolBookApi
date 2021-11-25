package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LessonMysql struct {
	db *sqlx.DB
}

func NewLessonMysql(db *sqlx.DB) *LessonMysql {
	return &LessonMysql{db: db}
}
func (r *LessonMysql) GetAll(userId, courseId int) ([]todo.Lesson,error){
	var lessons []todo.Lesson
	query:=fmt.Sprintf(`SELECT Idзанятия, занятия.название, занятия.описание , СтатусЗанятия , ДоступностьДляУченика , ДоступностьДляПреподавателя
                              FROM  %s INNER JOIN %s USING(Idкурса)
                              INNER JOIN %s USING(Idкурса) WHERE idКурса = ? AND idУченика = ?`, lessonTables, courseTables, learnerCoursesTables)
	if err:=r.db.Select(&lessons,query,courseId,userId);err!=nil{
		return nil,err
	}
	return lessons,nil
}
func (r *LessonMysql) GetById(userId, courseId ,lessonId int) (todo.Lesson,error){
	var lesson todo.Lesson
	query:=fmt.Sprintf(`SELECT Idзанятия, занятия.название, занятия.описание , СтатусЗанятия , ДоступностьДляУченика , ДоступностьДляПреподавателя
                              FROM  %s INNER JOIN %s USING(Idкурса)
                              INNER JOIN %s USING(Idкурса) WHERE idКурса = ? AND idУченика = ? AND IdЗанятия = ? AND ДоступностьДляУченика = 1`, lessonTables, courseTables, learnerCoursesTables)
	if err:=r.db.Get(&lesson,query,courseId,userId,lessonId);err!=nil{
		return todo.Lesson{},err
	}
	return lesson,nil
}