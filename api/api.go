package api

import (
	"fmt"
	"net/http"

	"github.com/douglastaylorb/api-students/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudent)
	api.Echo.GET("/students/:id", getStudent)
	api.Echo.PUT("/students/:id", updateStudent)
	api.Echo.DELETE("/students/:id", deleteStudent)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func getStudents(c echo.Context) error {
	students, err := db.GetStudents()

	if err != nil {
		return c.String(http.StatusNotFound, "Error to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func createStudent(c echo.Context) error {
	student := db.Student{}
	err := c.Bind(&student)
	if err != nil {
		return err
	}

	if err := db.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.String(http.StatusOK, "Student Created")
}

func getStudent(c echo.Context) error {
	id := c.Param("id")
	getStud := fmt.Sprintf("Get Student with ID: %s", id)
	return c.String(http.StatusOK, getStud)
}

func updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update: %s", id)
	return c.String(http.StatusOK, updateStud)
}

func deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deletetStud := fmt.Sprintf("Delete Student: %s", id)
	return c.String(http.StatusOK, deletetStud)
}
