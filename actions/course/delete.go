package course

import (
	"github.com/a-berahman/educative/repository"
)

//CourseAction ...
type DeleteCourseAction struct {
	Courser repository.Courser
	Id      int
}

//Execute prepares process of delete single course by courseId action
func (d *DeleteCourseAction) Execute() error {
	err := d.Courser.DeleteCourseByID(d.Id)
	if err != nil {
		return err
	}

	return nil
}
