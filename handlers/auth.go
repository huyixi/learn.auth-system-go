package handlers

import (
    "github.com/gin-gonic/gin"
    "auth-system-go/models"
    "auth-system-go/utils"
    "net/http"
    "auth-system-go/db"
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

    _, err = db.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, hashedPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert user into database"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var dbUser models.User
    err := db.DB.QueryRow("SELECT * FROM users WHERE username = ?", user.Username).Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not query user from database"})
        return
    }

    if !utils.CheckPasswordHash(user.Password, dbUser.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
        return
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetProfile(c *gin.Context) {
    userID, _ := c.Get("userID")
    
    var user models.User
    err := db.DB.QueryRow("SELECT * FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not query user from database"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}
