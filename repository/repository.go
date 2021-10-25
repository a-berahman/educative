package repository

import (
	"database/sql"

	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository/course"
	"github.com/a-berahman/educative/repository/student"
	"github.com/a-berahman/educative/repository/studentcourse"
)

// GetRepository returns repository  instace
// solution is implemented by Factory design pattern
func GetRepository(repositoryConst common.Repository, db *sql.DB) interface{} {
	switch repositoryConst {
	case common.Student:
		return student.NewStudent(db)
	case common.Course:
		return course.NewCourse(db)
	case common.StudentCourse:
		return studentcourse.NewStudentCourse(db)
	}

	return nil
}

//Studenter is implemented by the objects to promote student repository
type Studenter interface {
	Insert(model *models.Student) (int, error)
	GetAllStudents() ([]models.Student, error)
	UpdateStudentByID(id int, name string, email string, phone string) error
}

//Courser is implemented by the objects to promote course repository
type Courser interface {
	Insert(model *models.Course) (int, error)
	DeleteCourseByID(id int) error
}

//StudentCourser is implemented by the objects to promote studentCourse repository
type StudentCourser interface {
	Insert(model *models.StudentCourse) error
	GetAllCoursesByStudentId(id int) ([]models.StudentCourse, error)
}
