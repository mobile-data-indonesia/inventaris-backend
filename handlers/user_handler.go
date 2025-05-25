package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/mobile-data-indonesia/inventaris-backend/services"
	"github.com/mobile-data-indonesia/inventaris-backend/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{UserService: s}
}

func (ctrl *UserHandler) Register(c *gin.Context) {
	var input validators.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.UserService.RegisterUser(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func (ctrl *UserHandler) Login(c *gin.Context) {
	var input validators.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := ctrl.UserService.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("accessToken", accessToken, 3600, "/", "", false, true)
	c.SetCookie("refreshToken", refreshToken, 3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "login successful"})
}

func (ctrl *UserHandler) UpdateUser(c *gin.Context) {
	log.Println("UpdateUser called")
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse multipart form"})
		return
	}
	for key, values := range c.Request.MultipartForm.Value {
		for _, v := range values {
			log.Printf("Form field: %s, Value: %s\n", key, v)
		}
	}
	for key, files := range c.Request.MultipartForm.File {
		for _, f := range files {
			log.Printf("File field: %s, Filename: %s\n", key, f.Filename)
		}
	}

	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var input validators.UpdateUserRequest
	if err := c.ShouldBindWith(&input, binding.FormMultipart); err != nil {
		log.Println("Form bind error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var profileImageUrl *string
	file, err := c.FormFile("profile_picture")
	if err == nil {
		// Buat nama file unik
		ext := filepath.Ext(file.Filename)
		fileName := userID.String() + ext
		dst := filepath.Join("uploads/users", fileName)

		if err := os.MkdirAll("uploads/users", os.ModePerm); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create upload directory"})
			return
		}

		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save profile picture"})
			return
		}

		profileImage := "uploads/users/" + fileName
		profileImageUrl = &profileImage
	}

	if err := ctrl.UserService.UpdateUser(userID, input, profileImageUrl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully"})
}

func (ctrl *UserHandler) GetAllUsers(c *gin.Context) {
	// log.Println("GetAllUsers called")
	users, err := ctrl.UserService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (ctrl *UserHandler) GetUserByID(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := ctrl.UserService.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (ctrl *UserHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "refresh token not found"})
		return
	}

	newAccessToken, err := ctrl.UserService.RefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("accessToken", newAccessToken, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "token refreshed successfully"})
}
