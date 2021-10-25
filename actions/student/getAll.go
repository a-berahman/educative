package student

import (
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
)

//GetAllAction ...
type GetAllAction struct {
	Studenter repository.Studenter
}

// Execute prepares process of getting user list action
func (g *GetAllAction) Execute() ([]models.Student, error) {
	res, err := g.Studenter.GetAllStudents()
	if err != nil {
		return nil, err
	}
	return res, nil
}
