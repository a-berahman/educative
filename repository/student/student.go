package student

import (
	"database/sql"
	"fmt"

	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/util"

	"go.uber.org/zap"
)

const (
	createStudentQuery  = `INSERT INTO student (name, email, phone) VALUES ($1, $2, $3) RETURNING id`
	getAllStudentsQuery = `SELECT student.id, student.name, student.email, student.phone,(SELECT name as courseTitle FROM course WHERE ID = studentcourse.courseid)
						   FROM student INNER JOIN studentCourse on studentCourse.StudentId = student.id`
	updateStudentByID = `UPDATE student SET name=$1, email=$2, phone=$3 WHERE id=$4`
)

//Repository presents basic property for Student repository
type Repository struct {
	db  *sql.DB
	log *zap.SugaredLogger
}

//New returns new instance of Student repository
func NewStudent(db *sql.DB) *Repository {
	return &Repository{
		db:  db,
		log: util.Logger(),
	}
}

//Insert adds a new record in in student table
func (r *Repository) Insert(model *models.Student) (int, error) {
	var id int
	err := r.db.QueryRow(createStudentQuery, model.Name, model.Email, model.Phone).Scan(&id)
	if err != nil {
		r.log.Errorw("faild to insert in student table",
			"error", err,
			"model", model,
		)
		return id, err
	}
	model.ID = id
	return id, nil

}

//GetAllStudents returns all students
func (r *Repository) GetAllStudents() ([]models.Student, error) {
	var studentList []models.Student
	items, err := r.db.Query(getAllStudentsQuery)
	if err != nil {
		r.log.Errorw("failed to get all students",
			"error", err,
		)
		return nil, err
	}
	fmt.Println("items are : ", items)
	defer items.Close()
	for items.Next() {
		var currStudent models.Student
		err = items.Scan(&currStudent.ID, &currStudent.Name, &currStudent.Email, &currStudent.Phone, &currStudent.CourseTitle)
		if err != nil {
			return studentList, err
		}
		studentList = append(studentList, currStudent)
	}

	return studentList, nil
}

//UpdateStudentByID updates student by id
func (r *Repository) UpdateStudentByID(id int, name string, email string, phone string) error {
	res, err := r.db.Exec(updateStudentByID, name, email, phone, id)

	if err != nil {
		r.log.Errorw("failed to Update student ",
			"error", err,
			"res", res,
			"id", id,
		)
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		r.log.Errorw("Update student RowsAffected error",
			"error", err,
			"res", res,
			"id", id,
		)
		return err
	}

	return nil
}
