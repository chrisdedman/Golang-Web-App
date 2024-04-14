package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/internal/models"
	"github.com/golang-web-app/internal/utils"
)

// User represents a user in the database for deletion method.
// This struct is used to fully delete a user data from the database.
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

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
}

func (s *Server) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := s.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (s *Server) AuthenticateUser(email, password string) (string, error) {
	var user models.User
	if err := s.db.Model(models.User{}).Where("email = ?", email).Take(&user).Error; err != nil {
		return "", fmt.Errorf("account not found")
	}

	if err := utils.VerifyPassword(password, user.Password); err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return token, nil
}

func (s *Server) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (s *Server) DeleteUser(c *gin.Context) {
	user_id := c.Param("user_id")

	var user User
	if err := s.db.Where("id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := s.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func (s *Server) UpdateAccount(c *gin.Context) {
	userID := c.Param("user_id")

	var user models.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized to update account"})
		return
	}

	var input models.UpdateAccount
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Username != "" {
		user.Username = input.Username
	}

	if input.Password != "" {
		user.Password = input.Password
		utils.HashPassword(&user)
	}

	if err := s.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User account updated"})
}
