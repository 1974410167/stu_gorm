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
	DB.Model(student).Association("Teachers").Find(&teachers)
	fmt.Println(teachers[0])
}