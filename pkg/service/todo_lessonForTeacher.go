package service

import todo "FactorialSchoolBook"

func (s *lessonService) GetAllForTeacher(userId, courseId int) ([]todo.Lesson,error){
	return s.repo.GetAllForTeacher(userId, courseId)
}
func (s *lessonService) GetByIdForTeacher(userId, courseId, lessonId int) (todo.Lesson,error){
	return s.repo.GetByIdForTeacher(userId, courseId, lessonId)
}
func (s *lessonService) CreateLessonForTeacher(lesson todo.Lesson,courseId ,userId int) (int, error){
	_,err:=s.courseRepo.GetByIdForTeacher(userId,courseId)
	if err!= nil{
		// teacher not teaching this course
		return 0,err
	}
	return s.repo.CreateLessonForTeacher(lesson,courseId)
}
