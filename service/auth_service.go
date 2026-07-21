package service

import (
	"errors"
	"fmt"

	"github.com/Bhandavya345/Employee-Management/models"
	"github.com/Bhandavya345/Employee-Management/repository"
	"github.com/Bhandavya345/Employee-Management/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(user *models.User) error
	Login(email, password string) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

// Signup
func (s *authService) Signup(user *models.User) error {

	// Check if user already exists
	existingUser, _ := s.repo.GetUserByEmail(user.Email)

	if existingUser != nil {
		return errors.New("email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return s.repo.CreateUser(user)
}

// Login
func (s *authService) Login(email, password string) (string, error) {

	user, err := s.repo.GetUserByEmail(email)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password), // Hash from DB
		[]byte(password),      // Password entered by user
	)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT Token
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)

	fmt.Printf("Token generated for user ID: %d, Email: %s, Role: %s", user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
