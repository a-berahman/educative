package studentcourse

import (
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
)

type GetAction struct {
	StudentCourser repository.StudentCourser
	StudentId      int
}

// Execute prepares process of getting userCourses list action
func (g *GetAction) Execute() ([]models.StudentCourse, error) {
	res, err := g.StudentCourser.GetAllCoursesByStudentId(g.StudentId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
