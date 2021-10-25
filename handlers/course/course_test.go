package course

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/models"
	"github.com/a-berahman/educative/repository/course"
	_ "github.com/a-berahman/educative/util/testing"
	"github.com/labstack/echo/v4"
)

func TestInsertRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	courseHandler := NewCourse(currConfig)
	specs := []struct {
		name    string
		payload CourseRequest
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_insert",
			payload: CourseRequest{
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
		bodyByte, _ := json.Marshal(spec.payload)
		fmt.Println(string(bodyByte))
		got, err := doRequest(http.MethodPost, "/api/v1/course", nil, nil, courseHandler.InsertRequest, bytes.NewBuffer(bodyByte))
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v", spec.expect.status, got.StatusCode)
		}
	}

}

func TestDeleteRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	courseHandler := NewCourse(currConfig)
	sampleCourseId, _ := course.NewCourse(currConfig.PostgresConnection).Insert(&models.Course{Name: "testCourse", ProfessorName: "testProfessorName", Description: "testDescription"})
	specs := []struct {
		name     string
		courseId int
		expect   struct {
			err    error
			status int
		}
	}{
		{
			name:     "success_delete_course",
			courseId: sampleCourseId,
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
		got, err := doRequest(http.MethodDelete, "/api/v1/courses/:courseId", []string{"courseId"}, []string{strconv.Itoa(spec.courseId)}, courseHandler.DeleteCourseByIDRequest, nil)
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v", spec.expect.status, got.StatusCode)
		}
	}
}

func TestDeleteWithInvalidInputsRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	courseHandler := NewCourse(currConfig)
	specs := []struct {
		name     string
		courseId string
		expect   struct {
			err    error
			status int
		}
	}{
		{
			name:     "fail_delete_course_by_invalid_id",
			courseId: "0",
			expect: struct {
				err    error
				status int
			}{
				err:    nil,
				status: http.StatusOK,
			},
		},
		{
			name:     "fail_delete_course_by_null_id",
			courseId: "",
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
		got, err := doRequest(http.MethodDelete, "/api/v1/courses/:courseId", []string{"courseId"}, []string{spec.courseId}, courseHandler.DeleteCourseByIDRequest, nil)
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v", spec.expect.status, got.StatusCode)
		}
	}

}

func doRequest(methodType, baseURL string, paramsName, paramsValue []string, echoFunc func(c echo.Context) error, body *bytes.Buffer) (*http.Response, error) {

	req, err := http.NewRequest(methodType, "/", nil)
	if body != nil {
		req, err = http.NewRequest(methodType, "/", body)
	}

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)

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
