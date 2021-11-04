package service

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/repository"
	"mime/multipart"
)

type Authorization interface {

	CreateUser(user todo.User) (int,error)
	GenerateToken(username, password string) (string,error)
	ParseToken(token string) (todo.UserAuth,error)

	CreateUserForTeacher(user todo.User) (int,error)
	GenerateTokenForTeacher(username, password string) (string,error)
	ParseTokenForTeacher(token string) (todo.UserAuth,error)
}
type TodoCourse interface {
	GetAll(userId int) ([]todo.Course,error)
	GetById(userId, courseId int) (todo.Course,error)

	GetAllForTeacher(userId int) ([]todo.Course,error)
	GetByIdForTeacher(userId, courseId int) (todo.Course,error)
}
type TodoLesson interface {

	GetAll(userId,courseId int) ([]todo.Lesson,error)
	GetById(userId,courseId,lessonID int) (todo.Lesson,error)

	GetAllForTeacher(userId,courseId int) ([]todo.Lesson,error)
	GetByIdForTeacher(userId,courseId,lessonID int) (todo.Lesson,error)
	CreateLessonForTeacher(lesson todo.Lesson,courseId int,userId int) (int,error)
}
type TodoHomework interface {
	AddHomeworkForTeacher(courseIdStr , lessonIdStr string, mR *multipart.Reader) error
}
type Service struct{
	Authorization
	TodoLesson
	TodoCourse
	TodoHomework
}
func NewService(repos *repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoCourse: NewCourseService(repos.TodoCourse),
		TodoLesson: NewLessonService(repos.TodoLesson,repos.TodoCourse),
		TodoHomework: NewHomeworkService(repos.TodoLesson,repos.TodoCourse),
	}
}