package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sandbox-science/deep-focus/internal/database"
	"github.com/sandbox-science/deep-focus/internal/utils"
)

// The string "my_secret_key" is just an example and should be replaced with
// a secret key of sufficient length and complexity in a real-world scenario.
var jwtKey = []byte("my_secret_key")

// Login handles the login functionality.
// It receives a JSON payload containing the user's email and password,
// validates the credentials, generates a JWT token, and sets it as a cookie.
// If the login is successful, it returns a JSON response with a success message.
// If there are any errors during the login process, it returns a JSON response with an error message.
func Login(c *gin.Context) {

	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser database.User

	database.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &database.Claims{
		Role: existingUser.Role,
		Claims: jwt.MapClaims{
			"email": existingUser.Email,
			"exp":   expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in"})
}

// Signup handles the signup process for a user.
// It binds the JSON data from the request body to the `user` variable.
// If there is an error in binding the JSON data, it returns a JSON response with the error message.
// It checks if a user with the same email already exists in the database.
// If a user with the same email exists, it returns a JSON response with an error message.
// It generates a password hash for the user's password using the `GenerateHashPassword` function from the `utils` package.
// If there is an error in generating the password hash, it returns a JSON response with an error message.
// It creates a new user record in the database using the `Create` method from the `models.DB` object.
// Finally, it returns a JSON response with a success message indicating that the user has been created.
func Signup(c *gin.Context) {
	var user database.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser database.User

	database.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	database.DB.Create(&user)

	c.JSON(200, gin.H{"success": "user created"})
}

// Logout is a handler function that logs out the user by clearing the token cookie.
// It sets the token cookie value to an empty string and sets the expiration time to a negative value,
// effectively deleting the cookie from the client's browser.
// It then returns a JSON response indicating the success of the logout operation.
func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
