package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LessonInfoMysql struct {
	db *sqlx.DB
}

func NewLessonInfoMysql(db *sqlx.DB) *LessonInfoMysql {
	return &LessonInfoMysql{db: db}
}
func(r *LessonInfoMysql) GetHomeTask(lessonId int) (todo.HomeTask,error){
	var homeTask todo.HomeTask
	query:=fmt.Sprintf(`SELECT ДатаЗакрытияПубликации,ДатаЗакрытияПриема,ЯвляетсяВыложенным 
                               FROM %s
                               WHERE idЗанятия = ?`, homeTaskTables)
	err:=r.db.Get(&homeTask,query,lessonId)
	if err!=nil{
		return todo.HomeTask{},err
	}
	return homeTask,nil
}
func(r *LessonInfoMysql) GetHomework(userId, lessonId int) (todo.Homework,error){
	var homework todo.Homework
	query:=fmt.Sprintf(`SELECT Комментарий, Баллы1Задание, Баллы2Задание, ДопБаллы, ВсегоБаллов, СтатусОтправки, СтатусПроверки 
                               FROM %s INNER JOIN %s USING(УченикНаЗанятии)
                               WHERE idЗанятия = ? AND idУченика = ?`, homeworkTables, learnerInLessonTables)
	err:=r.db.Get(&homework,query,lessonId,userId)
	if err!=nil{
		return todo.Homework{},err
	}
	return homework,nil
}
func(r *LessonInfoMysql) PostHomework(userId, lessonId int) error{
	query:=fmt.Sprintf(`UPDATE %s INNER JOIN %s USING(УченикНаЗанятии)
                               INNER JOIN %s USING(idЗанятия)
                               INNER JOIN %s USING(idЗанятия)
                               SET  СтатусОтправки = true
                               WHERE idЗанятия = ? AND idУченика = ? AND ЯвляетсяВыложенным = true `, homeworkTables, learnerInLessonTables, lessonTables, homeTaskTables)
	_,err:=r.db.Exec(query,lessonId,userId)
	if err!=nil{
		return err
	}
	return nil
}
func(r *LessonInfoMysql) GetAttendance(userId, lessonId int) (todo.LearnerStatusAtLesson,error){
	var attendance todo.LearnerStatusAtLesson
	query:=fmt.Sprintf(`SELECT ПоднятаяРука, ПрисутствиеНаЗанятии 
                               FROM %s 
                               WHERE idЗанятия = ? AND idУченика = ?`, learnerInLessonTables)
	err:=r.db.Get(&attendance,query,lessonId,userId)
	if err!=nil{
		return todo.LearnerStatusAtLesson{},err
	}
	return attendance,nil
}
func(r *LessonInfoMysql) PostAttendance(userId, lessonId int) error{
	query:=fmt.Sprintf(`UPDATE %s INNER JOIN %s USING(idЗанятия)
                               SET  ПоднятаяРука= true
                               WHERE idЗанятия = ? AND idУченика = ? AND ВозможностьПоднятьРуку = true `, learnerInLessonTables, lessonTables)
	_,err:=r.db.Exec(query,lessonId,userId)
	if err!=nil{
		return err
	}
	return nil
}
