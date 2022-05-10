package main

import (
	"stu_gorm/Init_db"
	"stu_gorm/file1"
)


func main(){
	init_db.InitDB()
	// 迁移三个表
	DB := init_db.DB
	DB.AutoMigrate(&file1.Student{})
	DB.AutoMigrate(&file1.Teacher{})
	DB.AutoMigrate(&file1.School{})

	//file1.CreateSchool()
	//file1.CreateTeacher()
	//file1.CreateStudent()
	file1.SearchStudent()
}

