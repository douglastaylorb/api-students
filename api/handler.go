package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/douglastaylorb/api-students/schemas"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

// getStudents godoc
//
// @Summary 		Get all students
// @Description Get all students
// @Tags 				students
// @Accept 			json
// @Produce 		json
// @Param 			active query boolean false "Filter by active"
// @Success 		200 {object} schemas.StudentResponse
// @Failure 		404 {string} string "Error to get students"
// @Router 			/students/{id} [get]

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudents()

	if err != nil {
		return c.String(http.StatusNotFound, "Error to get students")
	}

	active := c.QueryParam("active")

	if active != "" {
		act, err := strconv.ParseBool(active)
		if err != nil {
			log.Error().Err(err).Msgf("[api] Error to parse boolean")
			return c.String(http.StatusInternalServerError, "Error to parse boolean")
		}
		students, err = api.DB.GetFilteredStudent(act)
	}

	ListOfStudents := map[string][]schemas.StudentResponse{"students": schemas.NewResponse(students)}

	return c.JSON(http.StatusOK, ListOfStudents)
}

func (api *API) createStudent(c echo.Context) error {
	studentReq := StudentRequest{}
	err := c.Bind(&studentReq)
	if err != nil {
		return err
	}

	if err := studentReq.Validate(); err != nil {
		log.Error().Err(err).Msgf("[api] Error to validate student")
		return c.String(http.StatusBadRequest, "Error to validating student")
	}

	student := schemas.Student{
		Name:   studentReq.Name,
		CPF:    studentReq.CPF,
		Email:  studentReq.Email,
		Age:    studentReq.Age,
		Active: *studentReq.Active,
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.JSON(http.StatusOK, student)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Invalid ID")
	}

	receivedStudent := schemas.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}

	updatingStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Error to get student")
	}

	student := updateStudentInfo(receivedStudent, updatingStudent)

	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to update student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) deleteStudent(c echo.Context) error {
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

	if err := api.DB.DeleteStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to delete student")
	}

	return c.JSON(http.StatusOK, student)

}

func updateStudentInfo(receivedStudent, student schemas.Student) schemas.Student {
	if receivedStudent.Name != "" {
		student.Name = receivedStudent.Name
	}

	if receivedStudent.CPF > 0 {
		student.CPF = receivedStudent.CPF
	}

	if receivedStudent.Email != "" {
		student.Email = receivedStudent.Email
	}

	if receivedStudent.Age > 0 {
		student.Age = receivedStudent.Age
	}

	if receivedStudent.Active != student.Active {
		student.Active = receivedStudent.Active
	}

	return student
}
