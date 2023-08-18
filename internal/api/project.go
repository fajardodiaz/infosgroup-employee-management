package api

import (
	"database/sql"
	"log"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type createProjectReq struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createProject(ctx *gin.Context) {
	var req createProjectReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	project, err := server.store.CreateProject(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, project)
}

type getProjectByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getProjectById(ctx *gin.Context) {
	var req getProjectByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	project, err := server.store.GetProject(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, project)
}

var paginationProjectParams = struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}{
	PageID:   1,
	PageSize: 20,
}

func (server *Server) getProjects(ctx *gin.Context) {
	if err := ctx.ShouldBindQuery(&paginationProjectParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.ListProjectsParams{
		Limit:  paginationProjectParams.PageSize,
		Offset: (paginationProjectParams.PageID - 1) * paginationProjectParams.PageSize,
	}

	projects, err := server.store.ListProjects(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, projects)
}

type deleteProjectRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProject(ctx *gin.Context) {
	var req deleteProjectRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := server.store.DeleteProject(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Project deleted successfully",
	})
}

type updateProjectIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateProjectNameRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateProject(ctx *gin.Context) {
	var reqName updateProjectNameRequest
	var reqId updateProjectIDRequest

	if err := ctx.ShouldBindJSON(&reqName); err != nil {
		log.Fatal("Error id")
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&reqId); err != nil {
		log.Fatal("Error id")
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	_, err := server.store.GetProject(ctx, reqId.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	args := database.UpdateProjectParams{
		ID:   reqId.ID,
		Name: reqName.Name,
	}

	updatedProject, err := server.store.UpdateProject(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedProject)

}
