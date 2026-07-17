package service

import (
	"github.com/Bhandavya345/Employee-Management/models"
	"github.com/Bhandavya345/Employee-Management/repository"
)

type EmployeeService interface {
	CreateEmployee(employee *models.Employee) error
	GetEmployees() ([]models.Employee, error)
	GetEmployeeByID(id uint) (*models.Employee, error)
	UpdateEmployee(employee *models.Employee) error
	DeleteEmployee(id uint) error

	GetEmployeesByDepartment(department string) ([]models.Employee, error)
	GetEmployeesBySalaryRange(min, max float64) ([]models.Employee, error)

	GetHighestSalaryEmployee() (*models.Employee, error)
	GetLowestSalaryEmployee() (*models.Employee, error)
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		repo: repo,
	}
}

// Create Employee
func (s *employeeService) CreateEmployee(employee *models.Employee) error {
	return s.repo.CreateEmployee(employee)
}

// Get All Employees
func (s *employeeService) GetEmployees() ([]models.Employee, error) {
	return s.repo.GetEmployees()
}

// Get Employee By ID
func (s *employeeService) GetEmployeeByID(id uint) (*models.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

// Update Employee
func (s *employeeService) UpdateEmployee(employee *models.Employee) error {
	return s.repo.UpdateEmployee(employee)
}

// Delete Employee
func (s *employeeService) DeleteEmployee(id uint) error {
	return s.repo.DeleteEmployee(id)
}

// Search By Department
func (s *employeeService) GetEmployeesByDepartment(department string) ([]models.Employee, error) {
	return s.repo.GetEmployeesByDepartment(department)
}

// Search By Salary Range
func (s *employeeService) GetEmployeesBySalaryRange(min, max float64) ([]models.Employee, error) {
	return s.repo.GetEmployeesBySalaryRange(min, max)
}

// Highest Salary Employee
func (s *employeeService) GetHighestSalaryEmployee() (*models.Employee, error) {
	return s.repo.GetHighestSalaryEmployee()
}

// Lowest Salary Employee
func (s *employeeService) GetLowestSalaryEmployee() (*models.Employee, error) {
	return s.repo.GetLowestSalaryEmployee()
}
