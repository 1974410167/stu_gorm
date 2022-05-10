package file1

import "stu_gorm/Init_db"

func CreateStudent(){
	DB := init_db.DB
	user1 := User{
		Name: "张三",
		Age: 16,
		Address: "糊涂大街道1",
	}
	user2 := User{
		Name: "李四",
		Age: 17,
		Address: "搜索街道1",
	}
	var teachers []*Teacher
	DB.Model(&Teacher{}).Where("id IN ?", []uint{1,2}).Find(&teachers)
	stu1 := &Student{
		User: user1,
		Score: 88,
		Class: 1,
		SchoolID: 6,
		Teachers: teachers,
	}
	stu2 := &Student{
		User: user2,
		Score: 99,
		Class: 2,
		SchoolID: 6,
		Teachers: teachers,
	}
	DB.Create([]*Student{stu1, stu2})

}