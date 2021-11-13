package helpers

import (
	"context"
	"os"

	"github.com/charanpy/todoapi/database"
	"github.com/charanpy/todoapi/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Protect(c *gin.Context) {


	token:=c.GetHeader("Authorization");

	if token=="" || len(token) == 0 {
		AppError(c,"UnAuthorized",401)
		return;
	}

	user,err:=ValidateToken(os.Getenv("SECRET"),token);


	if (err != nil || user.Id == "") {
		AppError(c,err.Error(),401)
		return;
	}

	var isUser model.User

	userId,_:= primitive.ObjectIDFromHex(user.Id)
	filter:=bson.D{{"_id",userId}}

	notfound := database.Client.Database("todoapp").Collection("users").FindOne(context.Background(),filter).Decode(&isUser)
	
	if (notfound != nil) {
		AppError(c,"Unauthorized",401)
		return
	}

	c.Set("id",user.Id)

	c.Next()
}