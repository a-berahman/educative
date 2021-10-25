package student

import (
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
)

//InsertAction ...
type InsertAction struct {
	Studenter repository.Studenter
	Student   *models.Student
}

// Execute prepares process of adding student action
func (i *InsertAction) Execute() error {

	_,err := i.Studenter.Insert(i.Student)
	if err != nil {
		return err
	}
	return nil
}
