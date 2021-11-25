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
	CreateLessonForTeacher(lesson todo.Lesson, courseId int) (int,error)

}

type TodoLessonInfo interface {

	GetHomeTask(lessonId int) (todo.HomeTask,error)
	GetHomework(userId, lessonId int) (todo.Homework,error)
	PostHomework(userId, lessonId int) error
	GetAttendance(userId, lessonId int) (todo.LearnerStatusAtLesson,error)
	PostAttendance(userId, lessonId int) error

	GetHomeTaskForTeacher(lessonId int) (todo.HomeTask,error)
	PostHomeTaskForTeacher(lessonId int) error
	GetAttendanceForTeacher(lessonId int) ([]todo.LearnerStatusAtLessonForTeacher,error)
	StartLessonForTeacher(courseId, lessonId, lessonNum int, publicationClosingDate,deliveryClosingDate string) error
	GetNumberLesson(lessonId int) (int,error)
	GetNumberNextLesson(courseId int) (int,error)
	FinishLessonForTeacher(lessonId int) error
	CheckHomeworkForTeacher( lessonId int,allHomeworks []todo.LearnerHomework) error
	CheckAttendanceForTeacher( lessonId int,Attendances []todo.LearnerAttendance) error

}
type Repository struct{
	Authorization
	TodoLesson
	TodoCourse
	TodoLessonInfo
}
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthMysql(db),
		TodoCourse:     NewCourseMysql(db),
		TodoLesson:     NewLessonMysql(db),
		TodoLessonInfo: NewLessonInfoMysql(db),
		}
}