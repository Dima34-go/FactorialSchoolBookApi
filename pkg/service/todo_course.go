package service

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/repository"
)

type CourseService struct{
	repo repository.TodoCourse
}
func NewCourseService(repo repository.TodoCourse) *CourseService {
	return &CourseService{
		repo: repo,
	}
}
func (s *CourseService) GetAll(id int) ([]todo.Course,error){
	return s.repo.GetAll(id)
}

func (s *CourseService) GetById(userId, courseId int) (todo.Course,error){
	return s.repo.GetById(userId, courseId)
}