package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/charanpy/todoapi/database"
	"github.com/charanpy/todoapi/helpers"
	"github.com/charanpy/todoapi/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddTodo(c *gin.Context) {
	var todo model.Todo


	userId,_:= primitive.ObjectIDFromHex(c.GetString("id"))

	if err := c.ShouldBindJSON(&todo); err != nil {
		helpers.AppError(c,err.Error(),400)
		return
	}

	todo.CreatedAt = time.Now()
	todo.UserId = userId

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
	sort:= bson.D{{"createdat",-1}}
	opt:= options.Find().SetSort(sort)
	userId,_:= primitive.ObjectIDFromHex(c.GetString("id"))

	filter:= bson.D{{"userId",userId}}

	cursor,err := database.Client.Database("todoapp").Collection("todos").Find(context.Background(),filter,opt)

	fmt.Println(cursor.Current)
	if err != nil {
		helpers.AppError(c,"No Todo Found",404)
		return
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


func DeleteTodo(c* gin.Context) {
	todoId,_:=primitive.ObjectIDFromHex(c.Param("todoId"))
	userId,_:=primitive.ObjectIDFromHex(c.GetString("id"))

	filter:=bson.D{{"_id",todoId},{"userId",userId}}

	result,err:= database.Client.Database("todoapp").Collection("todos").DeleteOne(context.Background(),filter)

	if err!=nil || result.DeletedCount == 0 {
		helpers.AppError(c,"Todo not found",401)
		return;
	}

	c.JSON(http.StatusOK,gin.H{
		"status":"success",
	})
}