package controller

import (
	"fmt"
	"net/http"

	"github.com/Bhandavya345/Employee-Management/models"
	"github.com/Bhandavya345/Employee-Management/service"
	"github.com/Bhandavya345/Employee-Management/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		Service: service,
	}
}

// POST /signup
func (ac *AuthController) Signup(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := ac.Service.Signup(&user)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "User Registered Successfully", nil)
}

// Login Request Model
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST /login
func (ac *AuthController) Login(c *gin.Context) {

	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := ac.Service.Login(request.Email, request.Password)

	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Login Successful", gin.H{
		"token": token,
	})
}

// GET /profile
func (ac *AuthController) GetProfile(c *gin.Context) {

	userIDValue, exists := c.Get("userID")
	fmt.Printf("User ID from context: %v, Exists: %v\n", userIDValue, exists)
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "User not found")
		return
	}

	userID := userIDValue.(uint)

	user, err := ac.Service.GetProfile(userID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Profile Retrieved Successfully", gin.H{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"role":    user.Role,
		"role_id": user.RoleID,
	})
}
