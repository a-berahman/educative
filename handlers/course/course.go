package course

import (
	//"educative/actions"	
	"net/http"

	"github.com/a-berahman/educative/actions/course"
	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
	"github.com/a-berahman/educative/util"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// CourseHandler is a HttpHandler that presents basic property for course
type CourseHandler struct {
	log        *zap.SugaredLogger
	CourseRepo repository.Courser
}

//CourseRequest represents course request schema
type CourseRequest struct {
	Name          string `json:"name"`
	ProfessorName string `json:"professorName"`
	Description   string `json:"description"`
}

// NewCourse creates course handler instance
func NewCourse(cfg *models.Configuration) *CourseHandler {
	return &CourseHandler{
		CourseRepo: repository.GetRepository(common.Course, cfg.PostgresConnection).(repository.Courser),
		log:        util.Logger()}
}

//InsertRequest prepares process of adding course request
func (u *CourseHandler) InsertRequest(c echo.Context) error {

	req := CourseRequest{}
	if err := c.Bind(&req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	obj := course.InsertAction{
		Courser: u.CourseRepo,
		Course:  &models.Course{Name: req.Name, ProfessorName: req.ProfessorName, Description: req.Description},
	}

	if err := obj.Execute(); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

//DeleteCourseByIDRequest prepares process of delete single course by courseId request
func (u *CourseHandler) DeleteCourseByIDRequest(c echo.Context) error {
	var courseId int
	if err := echo.PathParamsBinder(c).Int("courseId", &courseId).BindError(); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	obj := course.DeleteCourseAction{
		Courser: u.CourseRepo,
		Id:      courseId,
	}

	err := obj.Execute()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
