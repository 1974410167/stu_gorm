package file1

import (
	"fmt"
	"stu_gorm/Init_db"

)


func CreateSchool() {
	DB := init_db.DB
	school1 := &School{
		Address: "东路村",
		Name: "东路小学",
	}
	school2 := &School{
		Address: "方岗镇",
		Name: "方岗镇第一中学",
	}
	res := DB.Create([]*School{school1,school2})
	fmt.Println(res.RowsAffected)
	fmt.Println(school1.ID)
}