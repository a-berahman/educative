package student

import (
	"path/filepath"
	"testing"

	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
	_ "github.com/a-berahman/educative/util/testing"
)

func TestInsertAction(t *testing.T) {
	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)

	specs := []struct {
		student *models.Student
		expect  struct {
			err error
		}
	}{
		{
			student: &models.Student{
				Name:  "testStudent",
				Email: "testEmail",
				Phone: "testPhone",
			},
			expect: struct{ err error }{err: nil},
		},
	}
	for _, spec := range specs {
		obj := &InsertAction{
			Studenter: repository.GetRepository(common.Student, currConfig.PostgresConnection).(repository.Studenter),
			Student:   spec.student,
		}
		if got := obj.Execute(); got != spec.expect.err {
			t.Fatalf("expected got to be %v but it is %v", spec.expect.err, got)
		}
	}

}

func TestUpdateAction(t *testing.T) {
	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)

	specs := []struct {
		student *models.Student
		expect  struct {
			err error
		}
	}{
		{
			student: &models.Student{
				ID:    100,
				Name:  "modifiedStudent",
				Email: "modifiedEmail",
				Phone: "modfiedPhone",
			},
			expect: struct{ err error }{err: nil},
		},
	}
	for _, spec := range specs {
		obj := &UpdateAction{
			Studenter: repository.GetRepository(common.Student, currConfig.PostgresConnection).(repository.Studenter),
			Id:        spec.student.ID,
			Name:      spec.student.Name,
			Email:     spec.student.Email,
			Phone:     spec.student.Phone,
		}
		if got := obj.Execute(); got != spec.expect.err {
			t.Fatalf("expected got to be %v but it is %v", spec.expect.err, got)
		}
	}

}

func TestGetAllAction(t *testing.T) {
	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)

	specs := []struct {
		expect struct {
			err error
		}
	}{
		{
			expect: struct{ err error }{err: nil},
		},
	}
	for _, spec := range specs {
		obj := &GetAllAction{
			Studenter: repository.GetRepository(common.Student, currConfig.PostgresConnection).(repository.Studenter),
		}
		if _, got := obj.Execute(); got != spec.expect.err {
			t.Fatalf("expected got to be %v but it is %v", spec.expect.err, got)
		}
	}

}
