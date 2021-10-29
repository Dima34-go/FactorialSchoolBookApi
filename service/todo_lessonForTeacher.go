package service

import todo "FactorialSchoolBook"

func (s *lessonService) GetAllForTeacher(userId, courseId int) ([]todo.Lesson,error){
	return s.repo.GetAllForTeacher(userId, courseId)
}
func (s *lessonService) GetByIdForTeacher(userId, courseId, lessonId int) (todo.Lesson,error){
	return s.repo.GetByIdForTeacher(userId, courseId, lessonId)
}