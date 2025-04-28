package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/src/domain"
	"go_test/src/services"
	"net/http"
)

func UserSetup(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/users", GetUsers)
		api.GET("/users/:id", GetUserById)
		api.POST("/users", CreateUser)
		api.PATCH("/users/:id", PatchUser)

	}
}

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func PatchUser(c *gin.Context) {
	id := c.Param("id")
	var user domain.UserUpdateFields
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := services.PatchUser(id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
