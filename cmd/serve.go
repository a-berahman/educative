package cmd

import (
	"github.com/a-berahman/educative/common"
	"github.com/a-berahman/educative/config"
	"github.com/a-berahman/educative/handlers"
	"github.com/a-berahman/educative/util"

	//"educative/util"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run the application",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

// serve handles the serve command
func serve() {

	//load and initialize pre-requests
	currConfig := config.LoadConfig(configPath)
	util.Initialize()
	e := echo.New()
	api := e.Group("/api/v1")

	studentHandler := handlers.GetHandler(common.Student, currConfig).(handlers.Studenter)
	api.POST("/student", studentHandler.InsertRequest)
	api.GET("/student", studentHandler.GetAllStudentsRequest)
	api.PUT("/students/:studentId", studentHandler.UpdateStudentByIDRequest)

	courseHandler := handlers.GetHandler(common.Course, currConfig).(handlers.Courser)
	api.POST("/course", courseHandler.InsertRequest)
	api.DELETE("/courses/:courseId", courseHandler.DeleteCourseByIDRequest)

	studentCourseHandler := handlers.GetHandler(common.StudentCourse, currConfig).(handlers.StudentCourser)
	api.POST("/studentCourse", studentCourseHandler.InsertRequest)
	api.GET("/studentCourses/:studentId", studentCourseHandler.GetAllCoursesByStudentIdRequest)

	e.Use(allowOptionsRequests(), middleware.Recover())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.CFG.APP.Port)))

}

func init() {
	rootCMD.AddCommand(serveCmd)
}

func allowOptionsRequests() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
		AllowHeaders:     []string{"*"},
	})
}
