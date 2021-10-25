package studentcourse

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository/course"
	"github.com/a-berahman/educative/repository/student"
	"github.com/a-berahman/educative/repository/studentcourse"
	_ "github.com/a-berahman/educative/util/testing"
	"github.com/labstack/echo/v4"
)

func TestInsertRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentCourseHandler := NewStudentCourse(currConfig)

	sampleStudentId, _ := student.NewStudent(currConfig.PostgresConnection).Insert(&models.Student{Name: "testStudent", Email: "testEmail", Phone: "tesetPhone"})
	sampleCourseId, _ := course.NewCourse(currConfig.PostgresConnection).Insert(&models.Course{Name: "testCourse", ProfessorName: "testProfessorName", Description: "testDescription"})

	specs := []struct {
		name    string
		payload StudentCourseRequest
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_insert",
			payload: StudentCourseRequest{
				StudentId: sampleStudentId,
				CourseId:  sampleCourseId,
			},
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusOK,
			},
		},
		{
			name: "fail_insert_by_null_studentId",
			payload: StudentCourseRequest{
				CourseId: sampleCourseId,
			},
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusInternalServerError,
			},
		},
		{
			name: "fail_insert_by_null_courseId",
			payload: StudentCourseRequest{
				StudentId: sampleStudentId,
			},
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusInternalServerError,
			},
		},
		{
			name: "fail_insert_by_unavailable_courseId",
			payload: StudentCourseRequest{
				CourseId: 0,
			},
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusInternalServerError,
			},
		},
		{
			name: "fail_insert_by_unavailable_studentId",
			payload: StudentCourseRequest{
				StudentId: 0,
			},
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusInternalServerError,
			},
		},
	}

	for _, spec := range specs {
		bodyByte, _ := json.Marshal(spec.payload)

		got, err := doRequest(http.MethodPost, "/api/v1/studentCourse", nil, nil, nil,
			studentCourseHandler.InsertRequest, bytes.NewBuffer(bodyByte), "")
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v by: %v", spec.expect.status, got.StatusCode, spec.expect.err)
		}
	}
}

func TestInsertWithInvalidRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentCourseHandler := NewStudentCourse(currConfig)

	specs := []struct {
		name     string
		jsonBody string
		expect   struct {
			err    error
			status int
		}
	}{
		{
			name:     "fail_insert_by_string_inputs",
			jsonBody: "{\"studentId\":'test_studentId',\"courseId\":'test_courseId'}",
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusBadRequest,
			},
		},
	}

	for _, spec := range specs {
		got, err := doRequest(http.MethodPost, "/api/v1/studentCourse", nil, nil, nil,
			studentCourseHandler.InsertRequest, bytes.NewBuffer([]byte(spec.jsonBody)), "")

		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v by: %v", spec.expect.status, got.StatusCode, spec.expect.err)
		}
	}
}

func TestGetRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentCourseHandler := NewStudentCourse(currConfig)

	sampleStudentId, _ := student.NewStudent(currConfig.PostgresConnection).Insert(&models.Student{Name: "testStudent", Email: "testEmail", Phone: "tesetPhone"})
	sampleCourseId, _ := course.NewCourse(currConfig.PostgresConnection).Insert(&models.Course{Name: "testCourse", ProfessorName: "testProfessorName", Description: "testDescription"})
	studentcourse.NewStudentCourse(currConfig.PostgresConnection).Insert(&models.StudentCourse{StudentId: sampleStudentId, CourseId: sampleCourseId})

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
		got, err := doRequest(http.MethodGet, "/api/v1/studentCourses/:studentId", nil, []string{"studentId"}, []string{strconv.Itoa(sampleStudentId)}, studentCourseHandler.GetAllCoursesByStudentIdRequest, nil, "")
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v by: %v", spec.expect.status, got.StatusCode, spec.expect.err)
		}
	}

}

func doRequest(methodType, baseURL string, headers [][]string, paramsName, paramsValue []string, echoFunc func(c echo.Context) error, body *bytes.Buffer, jsonBody string) (*http.Response, error) {

	req, err := http.NewRequest(methodType, "/", nil)
	if body != nil {
		req, err = http.NewRequest(methodType, "/", body)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	c.SetPath(baseURL)
	c.SetParamNames(paramsName...)
	c.SetParamValues(paramsValue...)

	err = echoFunc(c)
	if err != nil {
		return nil, err
	}

	resp := rec.Result()
	defer resp.Body.Close()
	return resp, nil
}
