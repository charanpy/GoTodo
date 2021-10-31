package controller

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/charanpy/todoapi/database"
	"github.com/charanpy/todoapi/helpers"
	"github.com/charanpy/todoapi/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)



func SignUp(c *gin.Context) {
	var user model.User;


	if err:= c.ShouldBindJSON(&user); err != nil {
		helpers.AppError(c,err.Error(),400)
	
		return;
	}

	user.CreatedAt = time.Now()
	var err error;

	user.Password,err = helpers.HashPassword(user.Password);

	if err != nil {
		helpers.AppError(c,"Something went wrong",500)
		return
	}

	res,err:= database.Client.Database("todoapp").Collection("users").InsertOne(context.Background(),user)

	if err!=nil {
		helpers.AppError(c,"Something went wrong",500)
		return;
	}


	user.ID = res.InsertedID.(primitive.ObjectID)
	user.Password="";

	c.JSON(http.StatusOK,gin.H{
		"status":"success",
		"user":user,
		"insertedId":res.InsertedID,
	})

	
}

func Login(c *gin.Context) {
	var user model.User;

	if err:=c.ShouldBindJSON(&user);err!=nil {
				helpers.AppError(c,err.Error(),400)

	}

	// check user 

	var isUser model.User;
	filter:= bson.D{{"email",user.Email}}
	err:=database.Client.Database("todoapp").Collection("users").FindOne(context.Background(),filter).Decode(&isUser)

	if err != nil {
		helpers.AppError(c,"Invalid credentials",400)
		return
	}

	// check password

	if err != nil {
		helpers.AppError(c,"Something went wrong",500)
		return
	}

	isValidPassword:= helpers.ComparePassword(isUser.Password,user.Password)

	if !isValidPassword {
		helpers.AppError(c,"Invalid Credentials",400)
		return
	}


	token,err:= helpers.GenerateToken(isUser.ID.Hex(),os.Getenv("SECRET"))

	if err != nil {
		helpers.AppError(c,"Something went wrong",500)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":"success",
		"user":isUser,
		"token":token,
	})
}