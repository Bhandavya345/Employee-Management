package controller

import (
	"net/http"
	"strconv"

	"github.com/Bhandavya345/Employee-Management/models"
	"github.com/Bhandavya345/Employee-Management/service"
	"github.com/Bhandavya345/Employee-Management/utils"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		Service: service,
	}
}

// POST /employees
func (ec *EmployeeController) CreateEmployee(c *gin.Context) {

	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := ec.Service.CreateEmployee(&employee)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Employee Created Successfully", employee)
}

// GET /employees
func (ec *EmployeeController) GetEmployees(c *gin.Context) {

	employees, err := ec.Service.GetEmployees()

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employees Retrieved Successfully", employees)
}

// GET /employees/:id
func (ec *EmployeeController) GetEmployeeByID(c *gin.Context) {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	employee, err := ec.Service.GetEmployeeByID(uint(id))

	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employee Retrieved Successfully", employee)
}

// PUT /employees/:id
func (ec *EmployeeController) UpdateEmployee(c *gin.Context) {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	var employee models.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	employee.ID = uint(id)

	err := ec.Service.UpdateEmployee(&employee)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employee Updated Successfully", employee)
}

// DELETE /employees/:id
func (ec *EmployeeController) DeleteEmployee(c *gin.Context) {

	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := ec.Service.DeleteEmployee(uint(id))

	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employee Deleted Successfully", nil)
}

// GET /employees/search?department=IT
func (ec *EmployeeController) GetEmployeesByDepartment(c *gin.Context) {

	department := c.Query("department")

	employees, err := ec.Service.GetEmployeesByDepartment(department)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employees Retrieved Successfully", employees)
}

// GET /employees/salary?min=10000&max=50000
func (ec *EmployeeController) GetEmployeesBySalary(c *gin.Context) {

	min, _ := strconv.ParseFloat(c.Query("min"), 64)
	max, _ := strconv.ParseFloat(c.Query("max"), 64)

	employees, err := ec.Service.GetEmployeesBySalaryRange(min, max)

	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Employees Retrieved Successfully", employees)
}

// GET /employees/highest-salary
func (ec *EmployeeController) GetHighestSalary(c *gin.Context) {

	employee, err := ec.Service.GetHighestSalaryEmployee()

	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Highest Salary Employee", employee)
}

// GET /employees/lowest-salary
func (ec *EmployeeController) GetLowestSalary(c *gin.Context) {

	employee, err := ec.Service.GetLowestSalaryEmployee()

	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Lowest Salary Employee", employee)
}
