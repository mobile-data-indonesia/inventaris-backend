package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mobile-data-indonesia/inventaris-backend/models"
	"github.com/mobile-data-indonesia/inventaris-backend/validators"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) RegisterUser(input validators.RegisterRequest) error {
	var existing models.User
	if err := s.DB.Where("username = ?", input.Username).First(&existing).Error; err == nil {
		return errors.New("username already exists")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	user := models.User{
		UserID:          uuid.New(),
		Username:        input.Username,
		Password:        string(hashedPassword),
		Email:           input.Email,
		PhoneNumber:     input.PhoneNumber,
		Title:           input.Title,
		Role:            input.Role,
		Department:      input.Department,
	}

	return s.DB.Create(&user).Error
}

func (s *UserService) LoginUser(input validators.LoginRequest) (string, string, error) {
	var user models.User
	if err := s.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		return "", "", errors.New("invalid username or password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", "", errors.New("invalid username or password")
	}

	accessClaims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(15 * time.Minute).Unix(), //15 minutes
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":     time.Now().Add(30 * 24 * time.Hour).Unix(), // 30 days
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := s.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := s.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) UpdateUser(userID uuid.UUID, input validators.UpdateUserRequest, profileImageUrl *string) error {
	var user models.User
	if err := s.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		return err
	}

	user.Email = &input.Email
	user.PhoneNumber = &input.PhoneNumber
	user.Role = input.Role
	user.Title = input.Title
	user.Department = input.Department

	if profileImageUrl != nil {
		user.ProfileImageURL = *profileImageUrl
	}

	return s.DB.Save(&user).Error
}

func (s *UserService) RefreshToken(refreshToken string) (string, error) {
	secret := os.Getenv("JWT_REFRESH_SECRET")
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["username"] == nil || claims["role"] == nil {
		return "", errors.New("invalid refresh token claims")
	}

	accessClaims := jwt.MapClaims{
		"username": claims["username"],
		"role":     claims["role"],
		"exp":      time.Now().Add(15 * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}