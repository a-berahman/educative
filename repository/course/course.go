package course

import (
	"database/sql"
	"fmt"

	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository/studentcourse"
	"github.com/a-berahman/educative/util"
	"go.uber.org/zap"
)

const (
	createCourseQuery = `INSERT INTO course (name, professorName, description) VALUES ($1, $2, $3) RETURNING id`
	deleteCourseByID  = `DELETE FROM course WHERE id=$1`
)

//Repository presents basic property for Course repository
type Repository struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

//New returns new instance of Course repository
func NewCourse(db *sql.DB) *Repository {
	return &Repository{
		db:  db,
		log: util.Logger(),
	}
}

//Insert adds a new record in in course table
func (r *Repository) Insert(model *models.Course) (int, error) {
	var id int
	err := r.db.QueryRow(createCourseQuery, model.Name, model.ProfessorName, model.Description).Scan(&id)
	if err != nil {
		r.log.Errorw("failed to insert in course table",
			"error", err,
			"model", model,
		)
		return id, err
	}
	model.ID = id
	return id, nil
}

//DeleteCourseByID deletes course by id
func (r *Repository) DeleteCourseByID(id int) error {

	studentCourses, err := studentcourse.NewStudentCourse(r.db).GetAllCoursesByCoursetId(id)
	if len(studentCourses) > 0 {
		fmt.Println("delete failed")
		r.log.Errorw("failed to delete course which exists in studentCourse",
			"id", id,
		)
		return err
	}

	res, err := r.db.Exec(deleteCourseByID, id)
	if err != nil {
		r.log.Errorw("failed to delete course ",
			"error", err,
			"res", res,
			"id", id,
		)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		r.log.Errorw("delete course RowsAffected error",
			"error", err,
			"res", res,
			"id", id,
		)
		return err
	}
	return nil
}
