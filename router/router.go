package router

import (
	"net/http"

	"github.com/charanpy/todoapi/controller"
	"github.com/charanpy/todoapi/helpers"
	"github.com/gin-gonic/gin"
)

func SetRoutes() *gin.Engine {
	r := gin.Default()


	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.JSON(http.StatusInternalServerError,gin.H{
				"status":"error",
				"message": err,
			})
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	app:=r.Group("/api/v1/") 
	{
		app.POST("/todo",helpers.Protect,controller.AddTodo)
		app.GET("/todo",helpers.Protect, controller.GetTodos)
		app.POST("/register",controller.SignUp)
		app.POST("/login",controller.Login)
	}

	return r
}
