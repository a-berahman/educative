package student

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"github.com/a-berahman/educative/config"
	_ "github.com/a-berahman/educative/util/testing"
	"github.com/labstack/echo/v4"
)

func TestInsertRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentHandler := NewStudent(currConfig)
	specs := []struct {
		name    string
		payload StudentRequest
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_insert",
			payload: StudentRequest{
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
		bodyByte, _ := json.Marshal(spec.payload)
		got, err := doRequest(http.MethodPost, "/api/v1/student", nil, nil, nil, studentHandler.InsertRequest, bytes.NewBuffer(bodyByte))
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v", spec.expect.status, got.StatusCode)
		}
	}
}

func TestGetRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentHandler := NewStudent(currConfig)

	specs := []struct {
		name    string
		payload StudentRequest
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_get",
			payload: StudentRequest{},
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
		got, err := doRequest(http.MethodGet, "/api/v1/student", nil, nil, nil, studentHandler.GetAllStudentsRequest, nil)
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v", http.StatusOK, got.StatusCode)
		}
	}
}

func TestUpdateRequest(t *testing.T) {

	path, _ := filepath.Abs("./env.yaml")
	currConfig := config.LoadConfig(path)
	studentHandler := NewStudent(currConfig)
	specs := []struct {
		name    string
		payload StudentRequest
		expect  struct {
			err    error
			status int
		}
	}{
		{
			name: "success_update",
			payload: StudentRequest{
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
		bodyByte, _ := json.Marshal(spec.payload)
		got, err := doRequest(http.MethodPut, "/api/v1/students/:studentId", nil, []string{"studentId"}, []string{"1"}, studentHandler.UpdateStudentByIDRequest, bytes.NewBuffer(bodyByte))
		if err != nil {
			t.Fatalf("expected error to be nil, error: %v", err)
		}
		if got.StatusCode != spec.expect.status {
			t.Errorf("expected response status to be %v, got: %v", spec.expect.status, got.StatusCode)
		}
	}
}

func doRequest(methodType, baseURL string, headers [][]string, paramsName, paramsValue []string, echoFunc func(c echo.Context) error, body *bytes.Buffer) (*http.Response, error) {

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
