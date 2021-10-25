package models

//Course represents Course data model
type Course struct {
	ID            int    `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	ProfessorName string `db:"professorName" json:"professorName"`
	Description   string `db:"description" json:"description`
}
