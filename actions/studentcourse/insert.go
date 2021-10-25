package studentcourse

import (
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
)

type InsertAction struct {
	StudentCourser repository.StudentCourser
	StudentCourse  *models.StudentCourse
}

// Execute prepares process of adding studentCourse action
func (i *InsertAction) Execute() error {
	err := i.StudentCourser.Insert(i.StudentCourse)
	if err != nil {
		return err
	}
	return nil
}
