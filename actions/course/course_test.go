package course

import (
	"path/filepath"
	"testing"

	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository"
	_ "github.com/a-berahman/educative/util/testing"
)

func TestDeleteAction(t *testing.T) {
	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)

	specs := []struct {
		id     int
		expect struct {
			err error
		}
	}{
		{id: 999999, expect: struct{ err error }{err: nil}},
	}
	for _, spec := range specs {
		obj := &DeleteCourseAction{
			Courser: repository.GetRepository(common.Course, currConfig.PostgresConnection).(repository.Courser),
			Id:      9999999,
		}
		if got := obj.Execute(); got != spec.expect.err {
			t.Fatalf("expected got to be %v but it is %v", spec.expect.err, got)
		}
	}
}

func TestInsertAction(t *testing.T) {
	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)

	specs := []struct {
		course *models.Course
		expect struct {
			err error
		}
	}{
		{
			course: &models.Course{
				Name:          "testName",
				ProfessorName: "professonTest",
				Description:   "descTest",
			},
			expect: struct{ err error }{err: nil},
		},
	}
	for _, spec := range specs {
		obj := &InsertAction{
			Courser: repository.GetRepository(common.Course, currConfig.PostgresConnection).(repository.Courser),
			Course:  spec.course,
		}
		if got := obj.Execute(); got != spec.expect.err {
			t.Fatalf("expected got to be %v but it is %v", spec.expect.err, got)
		}
	}
}
