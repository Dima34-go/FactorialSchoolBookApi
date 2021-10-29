package repository

import (
	todo "FactorialSchoolBook"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) (int,error)
	GetUser(username, password string) (todo.User,error)
	CreateUserForTeacher(user todo.User) (int,error)
	GetUserForTeacher(username, password string) (todo.Teacher,error)
}
type TodoCourse interface {
	GetAll(id int) ([]todo.Course,error)
	GetById(userId, courseId int) (todo.Course,error)
	GetAllForTeacher(id int) ([]todo.Course,error)
	GetByIdForTeacher(userId, courseId int) (todo.Course,error)
}
type TodoLesson interface {
	GetAll(userId, courseId int) ([]todo.Lesson,error)
	GetById(userId, courseId ,lessonId int) (todo.Lesson,error)
	GetAllForTeacher(userId, courseId int) ([]todo.Lesson,error)
	GetByIdForTeacher(userId, courseId ,lessonId int) (todo.Lesson,error)
}
type Repository struct{
	Authorization
	TodoLesson
	TodoCourse
}
func NewRepository(db *sqlx.DB) *Repository{
	return &Repository{
		Authorization: NewAuthMysql(db),
		TodoCourse: NewCourseMysql(db),
		TodoLesson: NewLessonMysql(db),
		}
}