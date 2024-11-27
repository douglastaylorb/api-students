package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/douglastaylorb/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudent)
	e.GET("/students/:id", getStudent)
	e.PUT("/students/:id", updateStudent)
	e.DELETE("/students/:id", deleteStudent)

	// Start server
	if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start server", "error", err)
	}
}

// Handler
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
