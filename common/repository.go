package common

//Repository represents repository enum type
type Repository int

const (
	//Order const/enum for repository list
	Course Repository = iota + 1
	Student
	StudentCourse
)
