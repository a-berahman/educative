package student

import (
	"net/http"
	"path/filepath"
	"testing"	
	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/models"
	_ "github.com/a-berahman/educative/util/testing"
)

func TestInsert(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentRepo := NewStudent(currConfig.PostgresConnection)
	specs := []struct {
		name    string
		payload models.Student
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_insert",
			payload: models.Student{
				Name:  "testName",
				Email: "testEmail",
				Phone: "testPhone",
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
		_, err := studentRepo.Insert(&spec.payload)
		if err != spec.expect.err {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
	}
}

func TestGetAllStudents(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentRepo := NewStudent(currConfig.PostgresConnection)
	specs := []struct {
		name   string
		expect struct {
			err    error
			status int
		}
	}{
		{
			name: "success_get",
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
		list, err := studentRepo.GetAllStudents()
		if err != spec.expect.err {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if len(list) == 0 {
			t.Fatal("expected to list length not to be 0")
		}
	}
}

func TestUpdateStudentById(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentRepo := NewStudent(currConfig.PostgresConnection)
	studentModel := models.Student{Name: "testStudent", Email: "testEmail", Phone: "testPhone"}
	sampleId, _ := studentRepo.Insert(&studentModel)

	specs := []struct {
		name    string
		payload models.Student
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_update",
			payload: models.Student{
				Name:  "modifiedName",
				Email: "modifiedEmail",
				Phone: "modifiedPhone",
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
		err := studentRepo.UpdateStudentByID(sampleId, spec.payload.Name, spec.payload.Email, spec.payload.Phone)
		if err != spec.expect.err {
			t.Fatalf("expected error to be nil in UpdateStudentByID, error: %v", err)
		}
	}
}
