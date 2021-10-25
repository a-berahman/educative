package models

//StudentCourse represents StudentCourse data model
type StudentCourse struct {
	ID        int `db:"id" json:"id"`
	StudentId int `db:"studentId" json:"studentId"`
	CourseId  int `db:"courseId" json:"courseId"`
}
