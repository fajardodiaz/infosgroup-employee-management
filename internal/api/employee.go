package api

import (
	"database/sql"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

var paginationEmployeeParams = struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}{
	PageID:   1,
	PageSize: 20,
}

func (server *Server) getEmployees(ctx *gin.Context) {
	if err := ctx.ShouldBindQuery(&paginationEmployeeParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.ListEmployeesParams{
		Limit:  paginationEmployeeParams.PageSize,
		Offset: (paginationEmployeeParams.PageID - 1) * paginationEmployeeParams.PageSize,
	}

	employees, err := server.store.ListEmployees(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employees)
}

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
