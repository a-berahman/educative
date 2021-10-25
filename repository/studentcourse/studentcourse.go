package studentcourse

import (
	"database/sql"

	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/util"

	"go.uber.org/zap"
)

const (
	createStudentCourseQuery = `INSERT INTO studentCourse (studentId, courseId) VALUES ($1, $2) RETURNING id`
	getAllCoursesByStudentID = `SELECT id, studentId, courseId FROM studentCourse WHERE studentId=$1`
	getAllCoursesByCourseID  = `SELECT id, studentId, courseId FROM studentCourse WHERE courseId=$1`
)

//Repository presents basic property for studentCourse repository
type Repository struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

//New returns new instance of studentCourse repository
func NewStudentCourse(db *sql.DB) *Repository {
	return &Repository{
		db:  db,
		log: util.Logger(),
	}
}

//Insert adds a new record in in studentCourse table
func (r *Repository) Insert(model *models.StudentCourse) error {
	var id int
	err := r.db.QueryRow(createStudentCourseQuery, model.StudentId, model.CourseId).Scan(&id)
	if err != nil {
		r.log.Errorw("failed to insert in studentCourse table",
			"error", err,
			"model", model,
		)
		return err
	}
	model.ID = id
	return nil
}

//GetAllCoursesByStudentId returns all studentCourse by studentId
func (r *Repository) GetAllCoursesByStudentId(studentId int) ([]models.StudentCourse, error) {
	var studentCourseList []models.StudentCourse
	items, err := r.db.Query(getAllCoursesByStudentID, studentId)
	if err != nil {
		r.log.Errorw("failed to get all studentCourses by studentId",
			"error", err,
		)
		return nil, err
	}
	defer items.Close()
	for items.Next() {
		var currStudentCourse models.StudentCourse
		err = items.Scan(&currStudentCourse.ID, &currStudentCourse.StudentId, &currStudentCourse.CourseId)
		if err != nil {
			return studentCourseList, err
		}
		studentCourseList = append(studentCourseList, currStudentCourse)
	}

	return studentCourseList, nil
}

//GetAllCoursesByCoursesId returns all studentCourse by courseId
func (r *Repository) GetAllCoursesByCoursetId(courseId int) ([]models.StudentCourse, error) {
	var studentCourseList []models.StudentCourse
	items, err := r.db.Query(getAllCoursesByStudentID, courseId)
	if err != nil {
		r.log.Errorw("failed to get all studentCourses by courseId",
			"error", err,
		)
		return nil, err
	}
	defer items.Close()
	for items.Next() {
		var currStudentCourse models.StudentCourse
		err = items.Scan(&currStudentCourse.ID, &currStudentCourse.StudentId, &currStudentCourse.CourseId)
		if err != nil {
			return studentCourseList, err
		}
		studentCourseList = append(studentCourseList, currStudentCourse)
	}

	return studentCourseList, nil
}
