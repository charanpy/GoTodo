package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppError(c *gin.Context,message string,statusCode int) {

	status:= http.StatusInternalServerError;

	if(statusCode == 400) {
		status = http.StatusBadRequest
	} else if(statusCode == 404) {
		status = http.StatusNotFound
	} else {
		status = http.StatusUnauthorized
	}

	 c.AbortWithStatusJSON(status,gin.H{
		"status":"error",
		"message": message,
	})


}