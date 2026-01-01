package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang_restaurant_management/models"
)

// usersCollection will be set from main.go
var usersCollection *mongo.Collection

// SetUsersCollection allows injecting the collection from main.go
func SetUsersCollection(col *mongo.Collection) {
	usersCollection = col
}

// ErrorResponse used for Swagger failures
type ErrorResponse struct {
	Error string `json:"error"`
}

// ================= Swagger-Compatible Handler ==================

// @Summary Get all users
// @Description Returns a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} ErrorResponse
// @Router /users [get]
func GetUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error fetching users"})
		return
	}
	defer cursor.Close(ctx)

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error decoding users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
