package studentcourse

import (
	"net/http"

	"github.com/a-berahman/educative/actions/studentcourse"
	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
	"github.com/a-berahman/educative/util"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// StudentCourseHandler is a HttpHandler that presents basic property for studentCourse
type StudentCourseHandler struct {
	log               *zap.SugaredLogger
	StudentCourseRepo repository.StudentCourser
}

//StudentCourseRequest represents student course request schema
type StudentCourseRequest struct {
	StudentId int `json:"studentId" `
	CourseId  int `json:"courseId" `
}

// NewStudentCourse creates studentCourse handler instance
func NewStudentCourse(cfg *models.Configuration) *StudentCourseHandler {
	return &StudentCourseHandler{
		StudentCourseRepo: repository.GetRepository(common.StudentCourse, cfg.PostgresConnection).(repository.StudentCourser),
		log:               util.Logger()}
}

//InsertRequest prepares process of adding studentCourse request
func (u *StudentCourseHandler) InsertRequest(c echo.Context) error {
	req := new(StudentCourseRequest)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	obj := studentcourse.InsertAction{
		StudentCourser: u.StudentCourseRepo,
		StudentCourse:  &models.StudentCourse{StudentId: req.StudentId, CourseId: req.CourseId},
	}
	if err := obj.Execute(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

// GetAllCoursesByStudentIdRequest prepares process of getting userCourses list request
func (u *StudentCourseHandler) GetAllCoursesByStudentIdRequest(c echo.Context) error {

	var studentID int
	if err := echo.PathParamsBinder(c).Int("studentId", &studentID).BindError(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	obj := studentcourse.GetAction{
		StudentCourser: u.StudentCourseRepo,
		StudentId:      studentID,
	}
	res, err := obj.Execute()
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, res)
}
