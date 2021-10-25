//Package handlers is a bunch of handlers that support needed function for routes/api.go
package handlers

import (
	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/handlers/course"
	"github.com/a-berahman/educative/handlers/student"
	"github.com/a-berahman/educative/handlers/studentcourse"
	"github.com/a-berahman/educative/models"
	"github.com/labstack/echo/v4"
)

// GetHandler returns handler instace
// - solution is implemented by Factory design pattern
func GetHandler(repositoryConst common.Repository, db *models.Configuration) interface{} {
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

//Studenter is implemented by objects that promote student handler signatures
type Studenter interface {
	InsertRequest(c echo.Context) error
	GetAllStudentsRequest(c echo.Context) error
	UpdateStudentByIDRequest(c echo.Context) error
}

//Courser is implemented by objects that promote course handler signatures
type Courser interface {
	InsertRequest(c echo.Context) error
	DeleteCourseByIDRequest(c echo.Context) error
}

//StudentCourser is implemented by objects that promote student course handler signatures
type StudentCourser interface {
	InsertRequest(c echo.Context) error
	GetAllCoursesByStudentIdRequest(c echo.Context) error
}
