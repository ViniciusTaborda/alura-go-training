package main

import (
	"api_go_gin/controllers"
	"api_go_gin/database"
	"api_go_gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var Id int

func SetupTestingRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateNewStudentMock() {
	student := models.Student{
		Name: "Test Student",
		CPF:  "12427904910",
		RG:   "1338554812",
	}

	db := database.GetDbConnection()
	db.Create(&student)

	Id = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student

	db := database.GetDbConnection()
	db.Delete(&student, Id)
}

func TestGetSaluteShouldReturnOk(t *testing.T) {

	r := SetupTestingRoutes()

	r.GET("/:name", controllers.GetSalute)

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should return OK.")
}

func TestGetAllStudentsShouldReturnOk(t *testing.T) {
	CreateNewStudentMock()
	defer DeleteStudentMock()

	r := SetupTestingRoutes()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest(http.MethodGet, "/students", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should return OK.")
}

func TestGetStudentByCPFShouldReturnOk(t *testing.T) {
	CreateNewStudentMock()
	defer DeleteStudentMock()

	r := SetupTestingRoutes()
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)

	req, _ := http.NewRequest(http.MethodGet, "/students/cpf/12427904910", nil)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code, "Should return OK.")
}

func TestGetStudentByIdShouldReturnOk(t *testing.T) {
	CreateNewStudentMock()
	defer DeleteStudentMock()

	r := SetupTestingRoutes()
	r.GET("/students/:id", controllers.GetStudentById)

	req, _ := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/students/%s", strconv.Itoa(Id)),
		nil,
	)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var student models.Student

	json.Unmarshal(response.Body.Bytes(), &student)

	assert.Equal(t, http.StatusOK, response.Code, "Should return OK.")
	assert.Equal(t, "Test Student", student.Name, "Should return the same name.")
}

func TestDeleteStudentShouldReturnOk(t *testing.T) {
	CreateNewStudentMock()

	r := SetupTestingRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	req, _ := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("/students/%s", strconv.Itoa(Id)),
		nil,
	)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)
	assert.Equal(t, http.StatusOK, response.Code, "Should return OK.")
}

func TestEditStudentShouldReturnOk(t *testing.T) {
	CreateNewStudentMock()

	r := SetupTestingRoutes()
	r.PATCH("/students/:id", controllers.EditStudent)

	student := models.Student{
		Name: "Test Student",
		CPF:  "00427904910",
		RG:   "0038554812",
	}

	studentJson, _ := json.Marshal(student)

	req, _ := http.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/students/%s", strconv.Itoa(Id)),
		bytes.NewBuffer(studentJson),
	)

	response := httptest.NewRecorder()

	r.ServeHTTP(response, req)

	var studentResponse models.Student

	json.Unmarshal(response.Body.Bytes(), &studentResponse)

	assert.Equal(t, http.StatusOK, response.Code, "Should return OK.")
	assert.Equal(
		t,
		student.Name,
		studentResponse.Name,
		"Should return the same name.",
	)

}
