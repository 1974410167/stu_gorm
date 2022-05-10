package file1

// 用户基表
type User struct {
	ID uint
	Name string
	Age int
	Address string
	Gender string
}
// 学生表
type Student struct{
	User     User `gorm:"embedded"`
	Score    int
	Class    int
	Teachers []*Teacher `gorm:"many2many:student_teacher;"`
	SchoolID uint
}

// 老师表
type Teacher struct {
	User     User `gorm:"embedded"`
	Subject  string
	Salary   int
	Students []*Student `gorm:"many2many:student_teacher;"`
	SchoolID uint
}

// 学校表
type School struct {
	ID uint
	Address string
	Name string
	Teachers []Teacher
	Students []Student
}
