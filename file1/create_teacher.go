package file1

import "stu_gorm/Init_db"

func CreateTeacher(){
	DB := init_db.DB

	user1 := User{
		Name: "语文老师",
		Age: 23,
		Address: "糊涂大街道1",
	}
	user2 := User{
		Name: "数学老师",
		Age: 24,
		Address: "搜索街道1",
	}
	tea1 := &Teacher{
		User: user1,
		Subject: "语文",
		Salary: 7000,
		SchoolID: 6,
	}
	tea2 := &Teacher{
		User: user2,
		Subject: "数学",
		Salary: 8000,
		SchoolID: 7,
	}
	DB.Create([]*Teacher{tea1, tea2})

}
