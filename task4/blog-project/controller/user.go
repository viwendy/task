package controller

import (
	"blog-main/global"
	"blog-main/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "Invalid request body", nil)
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "Failed to create user", nil)
		return
	}
	var user model.User
	user.Username = req.Username
	user.Password = string(hashedPassword)
	db := global.GetDB()
	if err := db.Create(&user).Error; err != nil {
		global.ReturnJson(c, http.StatusInternalServerError, -1, "Failed to create user", nil)
		return
	}
	global.ReturnJson(c, http.StatusCreated, 0, "User registered successfully", nil)
}

func Login(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		global.ReturnJson(c, http.StatusBadRequest, -1, "Invalid request body", nil)
		return
	}
	db := global.GetDB()
	var storedUser model.User
	if err := db.Where("username = ?", req.Username).First(&storedUser).Error; err != nil {
		global.ReturnJson(c, http.StatusUnauthorized, -1, "Invalid username or password", nil)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password)); err != nil {
		global.ReturnJson(c, http.StatusUnauthorized, -1, "Invalid username or password", nil)
		return
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	global.ReturnJson(c, http.StatusOK, 0, "Login successful", gin.H{
		"token": tokenString,
		"userInfo": gin.H{
			"id":       storedUser.ID,
			"username": storedUser.Username,
		},
	})
}
