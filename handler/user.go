package handler

import (
	"net/http"
	"resto/entities"
	"time"

	"github.com/aidarkhanov/nanoid"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	helper "resto/helper"
)

func (h *Handler) InsertUser(c *gin.Context) {
	var userReq entities.User
	if err := c.BindJSON(&userReq); err != nil {
		panic(err)
	}
	// countEmail := fmt.Sprintf()
	userReq.User_id = nanoid.New()
	userReq.Created_at = time.Now().String()
	userReq.Updated_at = time.Now().String()
	err := h.Service.AddUser(userReq)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, gin.H{"created": userReq})
}
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Service.GetUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		c.IndentedJSON(http.StatusOK, user)
	}

}
func (h *Handler) GetUsers(c *gin.Context) {
	claims, err := helper.ValidateToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, map[string]interface{}{
			"success": false,
			"message": "Token required",
		})
		return
	}
	users, err := h.Service.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"users":  users,
		"userID": claims.ID,
	})
}
func (h *Handler) Login(c *gin.Context) {

	var loginReq entities.LoginReq

	if err := c.BindJSON(&loginReq); err != nil {
		panic(err)

	}

	userID, err := h.Service.GetLogin(loginReq.Email, loginReq.Password)
	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		panic(err)
	}
	token, expirationTime := helper.GenerateAllTokens(userID)
	http.SetCookie(
		c.Writer,
		&http.Cookie{
			Name:    "token",
			Value:   token,
			Expires: expirationTime,
		},
	)

	c.IndentedJSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func (h *Handler) Logout(c *gin.Context) {
	desCookie := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}

	http.SetCookie(
		c.Writer,
		&desCookie,
	)

	c.IndentedJSON(http.StatusOK, gin.H{
		"logout success": true,
	})
}

func (h *Handler) RefreshAccountToken(c *gin.Context) {
	claims, _ := helper.ValidateToken(c)

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(helper.JwtKey)

	if err != nil {
		c.Copy().Status(http.StatusInternalServerError)
	}
	http.SetCookie(
		c.Writer,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		},
	)

}
