package file1

import (
	"fmt"
	init_db "stu_gorm/Init_db"
)

func SearchStudent(){
	DB := init_db.DB
	// 查找关联，student和teacher为多对多，可以根据student的相关条件来找到teacher，反之亦然。
	var teachers []*Teacher
	var student *Student
	// id 为 1 student
	DB.Model(&Student{}).Where("id = ?", 1).Find(&student)
	// id为1的student对应的老师
	fmt.Println(student)
	var students3 []Student
	DB.Model(&Student{}).Preload("Teachers").Find(&students3)
	fmt.Printf("%+v", students3)
	DB.Model(student).Association("Teachers").Find(&teachers)
	fmt.Println(teachers[0])

	var school School
	//var students []*Student
	DB.Model(&School{}).Where("name = ?", "东路小学").First(&school)
	//DB.Model(&Student{}).Where("school_id = ?", school.ID).Find(&students)
	fmt.Println(school)

	var students1 []Student
	DB.Model(&school).Association("Students").Find(&students1)
	fmt.Println(students1)
}