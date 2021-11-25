package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
const(
	learnerTables = "ученики"
	learnerCoursesTables = "ученикЗаписанНаКурс"
	teachersTables = "преподаватели"
	teacherCoursesTables = "преподавательПреподаетКурс"
	courseTables ="курсы"
	lessonTables = "занятия"
	homeTaskTables = "ДомашнееЗадание"
	homeworkTables = "ДомашняяРабота"
	learnerInLessonTables = "УченикНаЗанятии"
)
type Config struct{
	Host string
	Port string
	Username string
	Password string
	DBName string
}
func NewMysqlDB(cfg Config)(*sqlx.DB,error){
	db,err := sqlx.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err!=nil{
		return nil,err
	}
	err= db.Ping()
	if err!=nil{
		return nil,err
	}
	return db,nil
}
