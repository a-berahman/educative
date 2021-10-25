package models

//Student represents Student data model
type Student struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Email       string `db:"email" json:"email"`
	Phone       string `db:"phone" json:"phone`
	CourseTitle string `db:"coursetitle" json:"coursetitle"`
}
