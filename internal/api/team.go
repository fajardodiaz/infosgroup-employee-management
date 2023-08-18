package api

import (
	"database/sql"
	"log"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type createTeamReq struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createTeam(ctx *gin.Context) {
	var req createTeamReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	team, err := server.store.CreateTeam(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, team)
}

type getTeamByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTeamById(ctx *gin.Context) {
	var req getTeamByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	team, err := server.store.GetTeam(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, team)
}

var paginationTeamParams = struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}{
	PageID:   1,
	PageSize: 20,
}

func (server *Server) getTeams(ctx *gin.Context) {
	// var req getTeamsRequest

	if err := ctx.ShouldBindQuery(&paginationTeamParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.ListTeamsParams{
		Limit:  paginationTeamParams.PageSize,
		Offset: (paginationTeamParams.PageID - 1) * paginationTeamParams.PageSize,
	}

	teams, err := server.store.ListTeams(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, teams)
}

type deleteTeamRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTeam(ctx *gin.Context) {
	var req deleteTeamRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := server.store.DeleteTeam(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Team deleted successfully",
	})
}

type updateTeamIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateTeamNameRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateTeam(ctx *gin.Context) {
	var reqName updateTeamNameRequest
	var reqId updateTeamIDRequest

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

	_, err := server.store.GetTeam(ctx, reqId.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	args := database.UpdateTeamParams{
		ID:   reqId.ID,
		Name: reqName.Name,
	}

	updatedTeam, err := server.store.UpdateTeam(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedTeam)

}
