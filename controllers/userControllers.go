package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// ================= Swagger-Compatible Handler ==================

// @Summary Create a new user
// @Description Creates a new user in the system. Auto-generates ID, UUID, and timestamps.
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid request body"})
		return
	}

	// 1. Generate BaseModel fields
	base := models.NewBaseModel()

	// 2. Merge BaseModel into input user
	user := models.User{
		BaseModel: base,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
	}

	// 3. Insert user into MongoDB
	result, err := usersCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	// 4. Set MongoDB ObjectID in the struct
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	// 5. Return created user
	c.JSON(http.StatusCreated, user)
}
