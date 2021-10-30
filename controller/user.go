package controller

import (
	"context"
	"fmt"
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
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":err.Error(),
		})
		return;
	}

	user.CreatedAt = time.Now()
	var err error;

	user.Password,err = helpers.HashPassword(user.Password);

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
				"status":"error",
			"message": "Something went wrong",
		})
		return
	}

	res,err:= database.Client.Database("todoapp").Collection("users").InsertOne(context.Background(),user)

	if err!=nil {
		c.JSON(http.StatusInternalServerError,gin.H{
				"status":"error",
			"message": "Something went wrong",
		})
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
		c.JSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"message":err.Error(),
		})
	}

	// check user 

	var isUser model.User;
	filter:= bson.D{{"email",user.Email}}
	err:=database.Client.Database("todoapp").Collection("users").FindOne(context.Background(),filter).Decode(&isUser)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":"Invalid credentials",
		})
		return
	}

	// check password

	if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"message":"Something went wrong",
		})
		return
	}

	isValidPassword:= helpers.ComparePassword(isUser.Password,user.Password)

	if !isValidPassword {
		c.JSON(http.StatusBadRequest,gin.H{
			"status":"error",
			"message":"Invalid credentials",
		})
		return
	}

	fmt.Println(isUser.ID, isUser.ID.Hex())

	token,err:= helpers.GenerateToken(isUser.ID.Hex(),os.Getenv("SECRET"))

	if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{
			"status":"error",
			"message":"Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":"success",
		"user":isUser,
		"token":token,
	})
}