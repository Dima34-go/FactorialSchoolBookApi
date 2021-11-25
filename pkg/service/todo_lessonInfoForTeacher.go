package service

import (
	todo "FactorialSchoolBook"
	"errors"
	"time"
)

func (s *lessonInfoService) GetHomeTaskForTeacher(userId, courseId, lessonId int) (todo.HomeTask,error){
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return todo.HomeTask{}, err
	}
	return s.lessonInfoRepo.GetHomeTaskForTeacher(lessonId)
}
func (s *lessonInfoService) PostHomeTaskForTeacher(userId, courseId, lessonId int) error{
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	return s.lessonInfoRepo.PostHomeTaskForTeacher(lessonId)
}
func (s *lessonInfoService) GetAttendanceForTeacher(userId,courseId,lessonId int) ([]todo.LearnerStatusAtLessonForTeacher,error){
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return nil,err
	}
	return s.lessonInfoRepo.GetAttendanceForTeacher(lessonId)
}
func (s *lessonInfoService) StartLessonForTeacher(userId,courseId,lessonId int) error{
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	nextLesson,err:=s.lessonInfoRepo.GetNumberLesson(lessonId)
	if err!=nil{
		return err
	}
	nowLesson,err:=s.lessonInfoRepo.GetNumberNextLesson(courseId)
	if err!=nil{
		return err
	}
	if nowLesson!=nextLesson{
		return errors.New("this lesson cannot open, it is not next lesson in course")
	}
	//получение времени окончания приема и публикации ДЗ
	var PublicationClosingDate,DeliveryClosingDate time.Time
	DeliveryClosingDate=time.Now().Add(7*24*time.Hour)
	PublicationClosingDate=time.Now().Add(4*24*time.Hour)

	return s.lessonInfoRepo.StartLessonForTeacher(courseId,lessonId,nextLesson,
		PublicationClosingDate.Format("2006-01-02 15:04:05"),DeliveryClosingDate.Format("2006-01-02 15:04:05"))
}
func (s *lessonInfoService) FinishLessonForTeacher(userId,courseId,lessonId int) error{
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	return s.lessonInfoRepo.FinishLessonForTeacher(lessonId)
}
func (s *lessonInfoService) CheckHomeworkForTeacher(userId,courseId,lessonId int,allHomework []todo.LearnerHomework) error{
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	//создаем новый массив
	var allCheckHomeworks []todo.LearnerHomework
	for _,homework:=range allHomework{
		if !(homework.Scores1Task == 0 && homework.Scores2Task == 0 && homework.Comment == "Комментарий пока не оставлен"){
			allCheckHomeworks=append(allCheckHomeworks,homework)
		}
	}
	return s.lessonInfoRepo.CheckHomeworkForTeacher(lessonId,allCheckHomeworks)
}
func (s *lessonInfoService) CheckAttendanceForTeacher(userId,courseId,lessonId int,Attendances []todo.LearnerAttendance) error{
	_,err:=s.repo.GetByIdForTeacher(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	return s.lessonInfoRepo.CheckAttendanceForTeacher(lessonId,Attendances)
}