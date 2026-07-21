package routes

import (
	"github.com/Bhandavya345/Employee-Management/controller"
	"github.com/Bhandavya345/Employee-Management/middleware"
	"github.com/Bhandavya345/Employee-Management/repository"
	"github.com/Bhandavya345/Employee-Management/service"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// Repository
	employeeRepo := repository.NewEmployeeRepository()
	userRepo := repository.NewUserRepository()

	// Service
	employeeService := service.NewEmployeeService(employeeRepo)
	authService := service.NewAuthService(userRepo)

	// Controller
	employeeController := controller.NewEmployeeController(employeeService)
	authController := controller.NewAuthController(authService)

	// -------------------------
	// Public Routes
	// -------------------------

	router.POST("/signup", authController.Signup)
	router.POST("/login", authController.Login)

	// -------------------------
	// Protected Routes
	// -------------------------

	employee := router.Group("/employees")
	employee.Use(middleware.AuthMiddleware())

	{
		employee.POST("", employeeController.CreateEmployee)
		employee.GET("", employeeController.GetEmployees)
		employee.GET("/:id", employeeController.GetEmployeeByID)
		employee.PUT("/:id", employeeController.UpdateEmployee)
		employee.DELETE("/:id", employeeController.DeleteEmployee)

		employee.GET("/search", employeeController.GetEmployeesByDepartment)
		employee.GET("/salary", employeeController.GetEmployeesBySalary)

		employee.GET("/highest-salary", employeeController.GetHighestSalary)
		employee.GET("/lowest-salary", employeeController.GetLowestSalary)
	}
}
