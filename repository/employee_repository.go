package repository

import (
	"errors"

	"github.com/Bhandavya345/Employee-Management/database"
	"github.com/Bhandavya345/Employee-Management/models"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
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

type EmployeeRepo struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &EmployeeRepo{}
}

// Create Employee
func (r *EmployeeRepo) CreateEmployee(employee *models.Employee) error {
	return database.DB.Create(employee).Error
}

// Get All Employees
func (r *EmployeeRepo) GetEmployees() ([]models.Employee, error) {

	var employees []models.Employee

	err := database.DB.Find(&employees).Error

	return employees, err
}

// Get Employee By ID
func (r *EmployeeRepo) GetEmployeeByID(id uint) (*models.Employee, error) {

	var employee models.Employee

	err := database.DB.First(&employee, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("employee not found")
	}

	return &employee, err
}

// Update Employee
func (r *EmployeeRepo) UpdateEmployee(employee *models.Employee) error {
	return database.DB.Save(employee).Error
}

// Delete Employee
func (r *EmployeeRepo) DeleteEmployee(id uint) error {

	result := database.DB.Delete(&models.Employee{}, id)

	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}

	return result.Error
}

// Search By Department
func (r *EmployeeRepo) GetEmployeesByDepartment(department string) ([]models.Employee, error) {

	var employees []models.Employee

	err := database.DB.
		Where("department = ?", department).
		Find(&employees).Error

	return employees, err
}

// Search By Salary Range
func (r *EmployeeRepo) GetEmployeesBySalaryRange(min, max float64) ([]models.Employee, error) {

	var employees []models.Employee

	err := database.DB.
		Where("salary BETWEEN ? AND ?", min, max).
		Find(&employees).Error

	return employees, err
}

// Get Employee with Highest Salary
func (r *EmployeeRepo) GetHighestSalaryEmployee() (*models.Employee, error) {

	var employee models.Employee

	err := database.DB.
		Order("salary DESC").
		First(&employee).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("no employee found")
	}

	return &employee, err
}

// Get Employee with Lowest Salary
func (r *EmployeeRepo) GetLowestSalaryEmployee() (*models.Employee, error) {

	var employee models.Employee

	err := database.DB.
		Order("salary ASC").
		First(&employee).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("no employee found")
	}

	return &employee, err
}
