package repository

import (
	todo "FactorialSchoolBook"
	"fmt"
)

func(r *LessonInfoMysql) GetHomeTaskForTeacher(lessonId int) (todo.HomeTask,error){
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
func(r *LessonInfoMysql) PostHomeTaskForTeacher(lessonId int) error{
	query:=fmt.Sprintf(`UPDATE %s
                               SET  ЯвляетсяВыложенным = true
                               WHERE idЗанятия = ?`, homeTaskTables)
	_,err:=r.db.Exec(query,lessonId)
	if err!=nil{
		return err
	}
	return nil
}
func(r *LessonInfoMysql) GetAttendanceForTeacher(lessonId int) ([]todo.LearnerStatusAtLessonForTeacher,error){
	var allInfo []todo.LearnerStatusAtLessonForTeacher
	query:=fmt.Sprintf(`SELECT idУченика, Имя, Фамилия, ПрисутствиеНаЗанятии, ПоднятаяРука, Комментарий, 
                               Баллы1Задание, Баллы2Задание, ДопБаллы, ВсегоБаллов, СтатусОтправки, СтатусПроверки
                               FROM %s INNER JOIN %s USING(УченикНаЗанятии)
                               INNER JOIN %s USING(idУченика)
                               WHERE idЗанятия = ?
`, homeworkTables, learnerInLessonTables, learnerTables)
	err:=r.db.Select(&allInfo,query,lessonId)
	if err!=nil{
		return nil,err
	}
	return allInfo,nil
}
func(r *LessonInfoMysql) StartLessonForTeacher(courseId, lessonId, lessonNum int, publicationClosingDate,deliveryClosingDate string) error{
	//открыть новое занятие
	tx,err:=r.db.Begin()
	if err!=nil{
		return err
	}
	//установка сроков приема работ включение возможности поднять руку
	query:=fmt.Sprintf(`UPDATE %s INNER JOIN %s USING(idЗанятия)
                               SET ДатаЗакрытияПубликации = ? , ДатаЗакрытияПриема = ?, ВозможностьПоднятьРуку = true, СтатусЗанятия = 'Проведено'
                               WHERE idЗанятия = ?`, homeTaskTables, lessonTables)
	_,err=tx.Exec(query,publicationClosingDate,deliveryClosingDate,lessonId)
	if err!=nil{
		tx.Rollback()
		return err
	}
	//открытие следующего занятия для учеников и преподавателя
	query=fmt.Sprintf(`UPDATE %s 
                               SET ДоступностьДляУченика = true
                               WHERE НомерЗанятияКурса = ? AND IdКурса = ?`, lessonTables)
	_,err=tx.Exec(query,lessonNum+1,courseId)
	if err!=nil{
		tx.Rollback()
		return err
	}
	query=fmt.Sprintf(`UPDATE %s 
                               SET ДоступностьДляПреподавателя = true
                               WHERE НомерЗанятияКурса = ? AND IdКурса = ?`, lessonTables)
	_,err=tx.Exec(query,lessonNum+2,courseId)
	if err!=nil{
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
func(r *LessonInfoMysql) FinishLessonForTeacher(lessonId int) error{
	query:=fmt.Sprintf(`UPDATE %s
                               SET ВозможностьПоднятьРуку = false
                               WHERE idЗанятия = ?
`, lessonTables)
	_,err:=r.db.Exec(query,lessonId)
	if err!=nil{
		return err
	}
	return nil
}

func(r *LessonInfoMysql) CheckHomeworkForTeacher( lessonId int,allHomeworks []todo.LearnerHomework) error{
	tx,err:=r.db.Begin()
	if err!=nil{
		return  err
	}
	query:=fmt.Sprintf(`UPDATE %s INNER JOIN %s USING(УченикНаЗанятии)
                              SET Комментарий = ?, Баллы1Задание = ?, Баллы2Задание = ?, ДопБаллы = ?, ВсегоБаллов = ?, СтатусПроверки = true
                              WHERE idЗанятия = ? AND idУченика = ?`, homeworkTables, learnerInLessonTables)
	for _,homework :=range allHomeworks{
		_,err=tx.Exec(query,homework.Comment,homework.Scores1Task,homework.Scores2Task,homework.AddScores,
			homework.FullScores,lessonId,homework.LearnerId)
		if err!=nil{
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
func(r *LessonInfoMysql) CheckAttendanceForTeacher( lessonId int,Attendances []todo.LearnerAttendance) error{
	tx,err:=r.db.Begin()
	if err!=nil{
		return  err
	}
	query:=fmt.Sprintf(`UPDATE %s
                               SET ПрисутствиеНаЗанятии = ?
                               WHERE idЗанятия = ? AND idУченика = ?`, learnerInLessonTables)
	for _,attendance :=range Attendances{
		_,err=tx.Exec(query,attendance.AttendanceAtLesson,lessonId,attendance.LearnerId)
		if err!=nil{
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func(r *LessonInfoMysql) GetNumberNextLesson( courseId int) (int,error){
	var lesNum todo.NextLesson
	query:=fmt.Sprintf(`SELECT СледующееЗанятие
                               FROM %s
                               WHERE idКурса = ?`, courseTables)
	err:=r.db.Get(&lesNum,query,courseId)
	if err!=nil{
		return 0,err
	}
	return lesNum.Number,nil
}

func(r *LessonInfoMysql) GetNumberLesson( lessonId int) (int,error){
	var lesNum todo.NowLesson
	query:=fmt.Sprintf(`SELECT НомерЗанятияКурса
                               FROM %s
                               WHERE IdЗанятия = ? `, lessonTables)
	err :=r.db.Get(&lesNum,query,lessonId)
	if err!=nil{
		return 0,err
	}
	return lesNum.Number,nil
}
