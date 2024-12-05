package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/douglastaylorb/api-students/db"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()

	if err != nil {
		return c.String(http.StatusNotFound, "Error to get students")
	}

	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c echo.Context) error {
	student := db.Student{}
	err := c.Bind(&student)
	if err != nil {
		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.String(http.StatusOK, "Student Created")
}

func (api *API) getStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid ID")
	}

	student, err := api.DB.GetStudent(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error to get student")
	}

	return c.JSON(http.StatusOK, student)

}

func (api *API) updateStudent(c echo.Context) error {
	id := c.Param("id")
	updateStud := fmt.Sprintf("Update: %s", id)
	return c.String(http.StatusOK, updateStud)
}

func (api *API) deleteStudent(c echo.Context) error {
	id := c.Param("id")
	deletetStud := fmt.Sprintf("Delete Student: %s", id)
	return c.String(http.StatusOK, deletetStud)
}
