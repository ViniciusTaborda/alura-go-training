package routes

import (
	"api_go_gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/home", controllers.GetIndexPage)
	r.GET("/students", controllers.GetAllStudents)
	r.GET("/students/:id", controllers.GetStudentById)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)
	r.GET("/:name", controllers.GetSalute)
	r.POST("/students", controllers.CreateNewStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PATCH("/students/:id", controllers.EditStudent)

	r.NoRoute(controllers.GetNotFoundPage)

	r.Run()
}
