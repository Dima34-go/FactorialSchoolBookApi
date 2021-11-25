package service

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/repository"
)

type lessonService struct{
	repo       repository.TodoLesson
	courseRepo repository.TodoCourse
}

func NewLessonService(repo repository.TodoLesson, courseRepo repository.TodoCourse) *lessonService {
	return &lessonService{repo: repo, courseRepo: courseRepo}
}

func (s *lessonService) GetAll(userId, courseId int) ([]todo.Lesson,error){
	return s.repo.GetAll(userId, courseId)
}
func (s *lessonService) GetById(userId, courseId, lessonId int) (todo.Lesson,error){
	return s.repo.GetById(userId, courseId, lessonId)
}