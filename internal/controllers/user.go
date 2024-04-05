package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/config/models"
	"github.com/golang-web-app/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the database for deletion method.
// This struct is used to fully delete a user data from the database.
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Register creates a new user and add it to the database.
func (s *Server) Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password, Email: input.Email}
	utils.HashPassword(&user)

	if err := s.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
	fmt.Println("User created")
}

// LoginCheck checks if the user exists in the database.
func (s *Server) LoginCheck(email, password string) (string, error) {
	var err error

	user := models.User{}

	if err = s.db.Model(models.User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return "", err
	}

	err = utils.VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}

// Login logs in the user and returns a JWT token.
func (s *Server) Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Email: input.Email, Password: input.Password}
	token, err := s.LoginCheck(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email or password is not correct"})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
	fmt.Println("Login successful")
}

// Logout logs out the user, and deletes the JWT token.
func (s *Server) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	fmt.Println("Logout successful")
}

// DeleteUser deletes a user from the database using the provided ID.
// All data related to the user will be deleted.
func (s *Server) DeleteUser(c *gin.Context) {
	user_id := c.Param("user_id")

	var user User
	if err := s.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := s.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	s.Logout(c)
	fmt.Println("User", user_id, "deleted")
}
