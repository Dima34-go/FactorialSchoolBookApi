package main

import (
	todo "FactorialSchoolBook"
	handler "FactorialSchoolBook/pkg"
	"FactorialSchoolBook/repository"
	"FactorialSchoolBook/service"
	"log"
)

func main(){
	/*if err:=InitConfig();err!=nil{
		log.Fatalf("error initializing configs: %s",err.Error())
	}*/
	db,err:=repository.NewMysqlDB(repository.Config{
		Host: "localhost",
		Port: "3306",
		Username: "root",
		Password: "password",
		DBName: "factorialdb",
	})
	if err!=nil{
		log.Fatalf("error initializing db: %s",err.Error())
	}
	repos:= repository.NewRepository(db)
	services:= service.NewService(repos)
	handlers:= handler.NewHandler(services)
	srv:=new(todo.Server)
	if err:=srv.Run("8080",handlers.InitRoutes());err!=nil{/*viper.GetString("port")*/
		log.Fatalf("error running http server: %v",err)
	}
}
/*func InitConfig()error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}*/
