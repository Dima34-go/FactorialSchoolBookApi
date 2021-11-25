package service

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/repository"
)

type lessonInfoService struct{
	repo           repository.TodoLesson
	courseRepo     repository.TodoCourse
	lessonInfoRepo repository.TodoLessonInfo
}
func NewLessonInfoService(repo repository.TodoLesson, courseRepo repository.TodoCourse,lessonInfoRepo repository.TodoLessonInfo) *lessonInfoService {
	return &lessonInfoService{repo: repo, courseRepo: courseRepo,lessonInfoRepo: lessonInfoRepo}
}



func (s *lessonInfoService) GetHomeTask(userId, courseId, lessonId int) (todo.HomeTask,error){
	_,err:=s.repo.GetById(userId,courseId,lessonId)
	if err!=nil{
		return todo.HomeTask{}, err
	}
	return s.lessonInfoRepo.GetHomeTask(lessonId)
}
func (s *lessonInfoService) GetHomework(userId, courseId, lessonId int) (todo.Homework,error){
	_,err:=s.repo.GetById(userId,courseId,lessonId)
	if err!=nil{
		return todo.Homework{}, err
	}
	return s.lessonInfoRepo.GetHomework(userId,lessonId)
}
func (s *lessonInfoService) PostHomework(userId, courseId, lessonId int) error {
	_,err:=s.repo.GetById(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	return s.lessonInfoRepo.PostHomework(userId,lessonId)
}
func (s *lessonInfoService) GetAttendance(userId, courseId, lessonId int) (todo.LearnerStatusAtLesson,error){
	_,err:=s.repo.GetById(userId,courseId,lessonId)
	if err!=nil{
		return todo.LearnerStatusAtLesson{}, err
	}
	return s.lessonInfoRepo.GetAttendance(userId,lessonId)
}
func (s *lessonInfoService) PostAttendance(userId, courseId, lessonId int) error{
	_,err:=s.repo.GetById(userId,courseId,lessonId)
	if err!=nil{
		return err
	}
	return s.lessonInfoRepo.PostAttendance(userId,lessonId)
}
