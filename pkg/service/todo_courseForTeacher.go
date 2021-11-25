package service

import todo "FactorialSchoolBook"

func (s *CourseService) GetAllForTeacher(id int) ([]todo.Course,error){
	return s.repo.GetAllForTeacher(id)
}

func (s *CourseService) GetByIdForTeacher(userId, courseId int) (todo.Course,error){
	return s.repo.GetByIdForTeacher(userId, courseId)
}