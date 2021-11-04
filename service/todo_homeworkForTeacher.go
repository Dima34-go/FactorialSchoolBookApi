package service

import (
	"io"
	"mime/multipart"
	"os"
)

func (s *homeworkService) AddHomeworkForTeacher(courseIdStr , lessonIdStr string, mR *multipart.Reader) error{
	for {
		part, err := mR.NextPart()
		if err == io.EOF {
			break
		}
		if err!=nil && err!=io.EOF{
			return err
		}
		if part.FileName() == "" {
			continue
		}

		dst, err := os.Create("static"+"/"+courseIdStr+"/"+lessonIdStr+ "/" + part.FileName())
		if err != nil {
			return err
		}
		_,err = io.Copy(dst, part)
		if err != nil {
			return err
		}
	}
	return nil
}