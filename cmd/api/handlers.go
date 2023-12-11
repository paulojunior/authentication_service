package main

import (
	contract "authentication/contract"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

var requestPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary Authenticate user and generate JWT token
// @Description Authenticates a user based on the provided credentials and generates a JWT token.
// @Tags authentication
// @Accept json
// @Produce json
// @Param email body string true "User's email" format(email)
// @Param password body string true "User's password" format(password)
// @Success 202 {string} string "JWT token generated successfully"
// @Failure 400 {string} string "Invalid credentials"
// @Failure 500 {string} string "Internal error generating JWT token"
// @Router /authenticate [post]
func HandlerAuthenticate(c echo.Context, userService contract.UserService) error {
	if err := c.Bind(&requestPayload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := userService.FindByEmail(requestPayload.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid credentials")
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		return c.JSON(http.StatusBadRequest, "invalid credentials")
	}

	jwt, err := GenerateJWTToken(user.FirstName, 1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "error generating JWT token")
	}

	return c.JSON(http.StatusAccepted, jwt)
}

func GenerateJWTToken(user string, expiration int) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * time.Duration(expiration)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key := []byte(viper.GetString("jwt_secret"))

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
