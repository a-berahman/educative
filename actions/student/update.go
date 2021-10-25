package student

import (
	"github.com/a-berahman/educative/repository"
)

//InsertAction ...
type UpdateAction struct {
	Studenter repository.Studenter
	Id        int
	Name      string
	Email     string
	Phone     string
}

// Execute prepares process of update single student by studentId action
func (u *UpdateAction) Execute() error {
	err := u.Studenter.UpdateStudentByID(u.Id, u.Name, u.Email, u.Phone)
	if err != nil {
		return err
	}

	return nil
}
