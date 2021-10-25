package course

import (
	"net/http"
	"path/filepath"
	"testing"

	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/models"
	_ "github.com/a-berahman/educative/util/testing"
)

func TestInsertCourse(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	courseRepo := NewCourse(currConfig.PostgresConnection)
	specs := []struct {
		name    string
		payload models.Course
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_insert",
			payload: models.Course{
				Name:          "testName",
				ProfessorName: "testProfessorName",
				Description:   "testDescription",
			},
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusOK,
			},
		},
	}
	for _, spec := range specs {
		_, err := courseRepo.Insert(&spec.payload)
		if err != spec.expect.err {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
	}
}

func TestDeleteCourseById(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	courseRepo := NewCourse(currConfig.PostgresConnection)
	courseModel := models.Course{Name: "course_test", ProfessorName: "professorName_test", Description: "description_test"}
	courseId, err := courseRepo.Insert(&courseModel)
	specs := []struct {
		name   string
		expect struct {
			err    error
			status int
		}
	}{
		{
			name: "success_delete_course",
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusOK,
			},
		},
	}
	for _, spec := range specs {
		err = courseRepo.DeleteCourseByID(courseId)
		if err != spec.expect.err {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
	}
}