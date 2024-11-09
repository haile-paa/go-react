package controllers

import (
	"context"

	"github.com/haile-paa/go-react/models"
	"github.com/haile-paa/go-react/util"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	cursor, err := util.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error { // Capitalized for export
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return err
	}
	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}
	InsertResult, err := util.Collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID = InsertResult.InsertedID.(primitive.ObjectID)
	return c.Status(201).JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error { // Capitalized for export
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = util.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}

func DeleteTodo(c *fiber.Ctx) error { // Capitalized for export
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
	filter := bson.M{"_id": objectID}
	_, err = util.Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}
