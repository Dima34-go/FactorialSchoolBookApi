package service

import (
	todo "FactorialSchoolBook"
	"FactorialSchoolBook/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
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
type TodoLessonInfo interface {

	GetHomeTask(id, courseId, lessonId int) (todo.HomeTask,error)
	GetHomework(userId, courseId, lessonId int) (todo.Homework,error)
	PostHomework(userId, courseId, lessonId int) error
	GetAttendance(userId, courseId, lessonId int) (todo.LearnerStatusAtLesson,error)
	PostAttendance(userId, courseId, lessonId int) error

	GetHomeTaskForTeacher(id, courseId, lessonId int) (todo.HomeTask,error)
	PostHomeTaskForTeacher(userId, courseId, lessonId int) error
	GetAttendanceForTeacher(userId,courseId,lessonId int) ([]todo.LearnerStatusAtLessonForTeacher,error)
	StartLessonForTeacher(userId,courseId,lessonId int) error
	FinishLessonForTeacher(userId,courseId,lessonId int) error
	CheckHomeworkForTeacher(userId,courseId,lessonId int,allHomework []todo.LearnerHomework) error
	CheckAttendanceForTeacher(userId,courseId,lessonId int,Attendances []todo.LearnerAttendance) error

}
type Service struct{
	Authorization
	TodoLesson
	TodoCourse
	TodoLessonInfo
}
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		TodoCourse:     NewCourseService(repos.TodoCourse),
		TodoLesson:     NewLessonService(repos.TodoLesson,repos.TodoCourse),
		TodoLessonInfo: NewLessonInfoService(repos.TodoLesson,repos.TodoCourse,repos.TodoLessonInfo),
	}
}