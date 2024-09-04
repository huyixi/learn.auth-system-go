package handlers

import (
    "github.com/gin-gonic/gin"
    "auth-system-go/models"
    "auth-system-go/utils"
    "net/http"
)

func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }
    user.Password = hashedPassword

    // 这里应该将用户保存到数据库

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 这里应该从数据库获取用户并验证密码

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetProfile(c *gin.Context) {
    userID, _ := c.Get("userID")
    c.JSON(http.StatusOK, gin.H{"message": "Profile accessed", "user_id": userID})
}
