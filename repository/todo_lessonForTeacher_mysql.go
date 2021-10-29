package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
)

func (r *LessonMysql) GetAllForTeacher(userId, courseId int) ([]todo.Lesson,error){
	var lessons []todo.Lesson
	query:=fmt.Sprintf(`SELECT Idзанятия, занятия.название, занятия.описание 
                              FROM  %s INNER JOIN %s USING(Idкурса)
                              INNER JOIN %s USING(Idкурса) WHERE idКурса = ? AND idПреподавателя = ?`,lessonTables,courseTables,teacherCoursesTables)
	if err:=r.db.Select(&lessons,query,courseId,userId);err!=nil{
		return nil,err
	}
	return lessons,nil
}
func (r *LessonMysql) GetByIdForTeacher(userId, courseId ,lessonId int) (todo.Lesson,error){
	var lesson todo.Lesson
	query:=fmt.Sprintf(`SELECT Idзанятия, занятия.название, занятия.описание 
                              FROM  %s INNER JOIN %s USING(Idкурса)
                              INNER JOIN %s USING(Idкурса) WHERE idКурса = ? AND idПреподавателя = ? AND IdЗанятия = ?`,lessonTables,courseTables,teacherCoursesTables)
	if err:=r.db.Get(&lesson,query,courseId,userId,lessonId);err!=nil{
		return todo.Lesson{},err
	}
	return lesson,nil
}