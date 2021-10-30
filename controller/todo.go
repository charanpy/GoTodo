package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/charanpy/todoapi/database"
	"github.com/charanpy/todoapi/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTodo(c *gin.Context) {
	var todo model.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	todo.CreatedAt = time.Now()

	result, err := database.Client.Database("todoapp").Collection("todos").InsertOne(context.Background(), todo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n %v", result, result)

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"todo":    todo,
		"_id":     result.InsertedID,
	})
}

func GetTodos(c *gin.Context) {
	// sort:= bson.D{{"createdat",-1}}
	// opt:=options.Find().SetSort(sort)
	filter:= bson.D{}

	cursor,err := database.Client.Database("todoapp").Collection("todos").Find(context.Background(),filter)

	fmt.Println(cursor.Current)
	if err != nil {
		log.Fatal(err)
	}

	var todos []primitive.M;

	for(cursor.Next(context.Background())) {
		var todo bson.M;
		err:= cursor.Decode(&todo)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	defer cursor.Close(context.Background())

	c.JSON(http.StatusOK,gin.H{
		"message":"success",
		"todos": todos,
	})
}