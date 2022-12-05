package controllers

import (
	"api_go_gin/database"
	"api_go_gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {

	var students []models.Student

	db := database.GetDbConnection()

	db.Find(&students)

	c.JSON(http.StatusOK, students)

}

func GetStudentById(c *gin.Context) {

	id := c.Params.ByName("id")

	var student models.Student

	db := database.GetDbConnection()
	db.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
	} else {
		c.JSON(http.StatusOK, student)
	}

}

func GetStudentByCPF(c *gin.Context) {

	cpf := c.Param("cpf")

	var student models.Student

	student.CPF = cpf

	db := database.GetDbConnection()
	db.Where(&student).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
	} else {
		c.JSON(http.StatusOK, student)
	}

}

func EditStudent(c *gin.Context) {

	id := c.Params.ByName("id")

	var student models.Student

	db := database.GetDbConnection()
	db.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		if err := models.ValidateStudentData(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db.Model(&student).UpdateColumns(student)
		c.JSON(http.StatusOK, student)
	}

}

func DeleteStudent(c *gin.Context) {

	id := c.Params.ByName("id")

	var student models.Student

	db := database.GetDbConnection()
	db.Delete(&student, id)

	c.JSON(http.StatusOK, student)

}

func CreateNewStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {

		if err := models.ValidateStudentData(&student); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		db := database.GetDbConnection()
		db.Create(&student)
		c.JSON(http.StatusOK, student)
	}

}

func GetSalute(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hey %s!", name),
	})
}

func GetIndexPage(c *gin.Context) {

	var students []models.Student

	db := database.GetDbConnection()

	db.Find(&students)

	c.HTML(http.StatusOK, "index.html", students)
}

func GetNotFoundPage(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
