package api

import (
	"database/sql"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/fajardodiaz/infosgroup-employee-management/internal/utils"
	"github.com/gin-gonic/gin"
)

// List employees handler func
var paginationEmployeeParams = struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=200"`
}{
	PageID:   1,
	PageSize: 50,
}

func (server *Server) getEmployees(ctx *gin.Context) {
	if err := ctx.ShouldBindQuery(&paginationEmployeeParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.ListEmployeesWithRelParams{
		Limit:  paginationEmployeeParams.PageSize,
		Offset: (paginationEmployeeParams.PageID - 1) * paginationEmployeeParams.PageSize,
	}

	employees, err := server.store.ListEmployeesWithRel(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

// Get employee by id handler func
type getEmployeeByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEmployeeById(ctx *gin.Context) {
	var req getEmployeeByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	employee, err := server.store.GetEmployee(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// delete employee handler func
type deleteEmployeeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteEmployee(ctx *gin.Context) {
	var req deleteEmployeeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := server.store.DeleteEmployee(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Employee deleted succcesfully",
	})
}

// Create employee handler func
type createEmployeeRequest struct {
	EmployeeCod string `json:"employee_cod" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	Birth       string `json:"birth" binding:"required"`
	IngressDate string `json:"ingress_date" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Gender      string `json:"gender" binding:"required,oneof=M F O"`
	StateID     int    `json:"state_id" binding:"required"`
	TeamID      int    `json:"team_id" binding:"required"`
	PositionID  int    `json:"position_id" binding:"required"`
}

func (server *Server) createEmployee(ctx *gin.Context) {
	var req createEmployeeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	dateBirth := utils.ConvertToDate(req.Birth)
	ingressDate := utils.ConvertToDate(req.IngressDate)
	endEvaluationDate := utils.Add3MonthsToDate(ingressDate)

	args := database.CreateEmployeeParams{
		EmployeeCod:       req.EmployeeCod,
		FullName:          req.FullName,
		Birth:             dateBirth,
		IngressDate:       ingressDate,
		EndEvaluationDate: endEvaluationDate,
		Phone:             sql.NullString{String: req.Phone, Valid: true},
		Gender:            sql.NullString{String: req.Gender, Valid: true},
		StateID:           sql.NullInt32{Int32: int32(req.StateID), Valid: true},
		PositionID:        sql.NullInt32{Int32: int32(req.PositionID), Valid: true},
		TeamID:            sql.NullInt32{Int32: int32(req.TeamID), Valid: true},
	}

	employee, err := server.store.CreateEmployee(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}

// Update employee handler func
type updateEmployeeIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateEmployeeRequest struct {
	EmployeeCod string `json:"employee_cod" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	Birth       string `json:"birth" binding:"required"`
	IngressDate string `json:"ingress_date" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	Gender      string `json:"gender" binding:"required,oneof=M F O"`
	StateID     int    `json:"state_id" binding:"required"`
	TeamID      int    `json:"team_id" binding:"required"`
	PositionID  int    `json:"position_id" binding:"required"`
}

func (server *Server) updateEmployee(ctx *gin.Context) {
	var req updateEmployeeRequest
	var reqID updateEmployeeIDRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	dateBirth := utils.ConvertToDate(req.Birth)
	ingressDate := utils.ConvertToDate(req.IngressDate)
	endEvaluationDate := utils.ConvertToDate(req.IngressDate)

	args := database.UpdateEmployeeParams{
		ID:                reqID.ID,
		EmployeeCod:       req.EmployeeCod,
		FullName:          req.FullName,
		Birth:             dateBirth,
		IngressDate:       ingressDate,
		EndEvaluationDate: endEvaluationDate,
		Phone:             sql.NullString{String: req.Phone, Valid: true},
		Gender:            sql.NullString{String: req.Gender, Valid: true},
		StateID:           sql.NullInt32{Int32: int32(req.StateID), Valid: true},
		PositionID:        sql.NullInt32{Int32: int32(req.PositionID), Valid: true},
		TeamID:            sql.NullInt32{Int32: int32(req.TeamID), Valid: true},
	}

	employee, err := server.store.UpdateEmployee(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employee)
}
