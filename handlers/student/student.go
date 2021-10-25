package student

import (
	"net/http"

	"github.com/a-berahman/educative/actions/student"
	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
	"github.com/a-berahman/educative/util"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// StudentHandler is a HttpHandler that presents basic property for student
type StudentHandler struct {
	log         *zap.SugaredLogger
	StudentRepo repository.Studenter
}

//StudentRequest represents student request schema
type StudentRequest struct {
	Name  string `json:"name" `
	Email string `json:"email" `
	Phone string `json:"phone"`
}

//NewStudent creates student handler instance
func NewStudent(cfg *models.Configuration) *StudentHandler {
	sh := &StudentHandler{
		StudentRepo: repository.GetRepository(common.Student, cfg.PostgresConnection).(repository.Studenter),
		log:         util.Logger()}
	return sh
}

//InsertRequest prepares process of adding student request
func (u *StudentHandler) InsertRequest(c echo.Context) error {

	req := new(StudentRequest)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	obj := student.InsertAction{
		Studenter: u.StudentRepo,
		Student:   &models.Student{Name: req.Name, Email: req.Email, Phone: req.Phone},
	}

	if err := obj.Execute(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// GetAllStudentsRequest prepares process of getting user list request
func (u *StudentHandler) GetAllStudentsRequest(c echo.Context) error {
	obj := student.GetAllAction{
		Studenter: u.StudentRepo,
	}
	res, err := obj.Execute()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateStudentByIDRequest prepares process of update single student by studentId request
func (u *StudentHandler) UpdateStudentByIDRequest(c echo.Context) error {
	req := new(StudentRequest)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	var studentID int
	if err := echo.PathParamsBinder(c).Int("studentId", &studentID).BindError(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	obj := student.UpdateAction{
		Studenter: u.StudentRepo,
		Id:        studentID,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
	}
	if err := obj.Execute(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
