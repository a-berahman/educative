package course

import (
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
)

type InsertAction struct {
	Courser repository.Courser
	Course  *models.Course
}

//Execute prepares process of adding course action
func (i *InsertAction) Execute() error {

	_,err := i.Courser.Insert(i.Course)
	if err != nil {
		return err
	}
	return nil
}
