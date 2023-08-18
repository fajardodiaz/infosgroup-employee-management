package api

import (
	"database/sql"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type AddUserToProjectRequest struct {
	IDProject  int64 `uri:"id_project" binding:"required,min=1"`
	IdEmployee int64 `uri:"id_employee" binding:"required,min=1"`
}

func (server *Server) AddUserToProject(ctx *gin.Context) {
	var req AddUserToProjectRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.AssignEmployeeToProjectParams{
		ProjectID:  req.IDProject,
		EmployeeID: req.IdEmployee,
	}

	assignment, err := server.store.AssignEmployeeToProject(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, assignment)
}

type getEmployeeProjectsRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getEmployeeProjects(ctx *gin.Context) {
	var req getEmployeeProjectsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	employeeProjects, err := server.store.GetEmployeeProjects(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, employeeProjects)
}

type getProjectEmployeesRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProjectEmployees(ctx *gin.Context) {
	var req getProjectEmployeesRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	projectEmployees, err := server.store.GetProjectEmployees(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, projectEmployees)
}

type DeleteEmployeeFromAProjectRequest struct {
	IDProject  int64 `uri:"id_project" binding:"required,min=1"`
	IdEmployee int64 `uri:"id_employee" binding:"required,min=1"`
}

func (server *Server) deleteEmployeeFromAProject(ctx *gin.Context) {
	var req DeleteEmployeeFromAProjectRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.RemoveEmployeeProjectParams{
		EmployeeID: req.IdEmployee,
		ProjectID:  req.IDProject,
	}

	err := server.store.RemoveEmployeeProject(ctx, args)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Employee removed from project successfully",
	})
}
