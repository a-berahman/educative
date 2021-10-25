package studentcourse

import (
	"path/filepath"
	"testing"

	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/repository"
	_ "github.com/a-berahman/educative/util/testing"
)

func TestGetByStudentIdAction(t *testing.T) {
	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)

	specs := []struct {
		StudentId int
		expect    struct {
			err error
		}
	}{
		{
			StudentId: 1000,
			expect:    struct{ err error }{err: nil},
		},
	}
	for _, spec := range specs {
		obj := &GetAction{
			StudentCourser: repository.GetRepository(common.StudentCourse, currConfig.PostgresConnection).(repository.StudentCourser),
			StudentId:      spec.StudentId,
		}
		if _, got := obj.Execute(); got != spec.expect.err {
			t.Fatalf("expected got to be %v but it is %v", spec.expect.err, got)
		}
	}

}
