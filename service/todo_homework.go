package service

import "FactorialSchoolBook/repository"

type homeworkService struct{
	repo repository.TodoLesson
	courseRepo repository.TodoCourse
}
func NewHomeworkService(repo repository.TodoLesson, courseRepo repository.TodoCourse) *homeworkService {
	return &homeworkService{repo: repo, courseRepo: courseRepo}
}