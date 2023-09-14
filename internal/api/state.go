package api

import (
	"database/sql"
	"log"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type createStateReq struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createState(ctx *gin.Context) {
	var req createStateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	state, err := server.store.CreateState(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, state)
}

type getStateByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getStateById(ctx *gin.Context) {
	var req getStateByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	state, err := server.store.GetState(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, state)
}

type getStateByIdNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) getStateIdByName(ctx *gin.Context) {
	var req getStateByIdNameRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	state, err := server.store.GetStateIdByName(ctx, req.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, state)
}

var paginationStateParams = struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}{
	PageID:   1,
	PageSize: 20,
}

func (server *Server) getStates(ctx *gin.Context) {

	if err := ctx.ShouldBindQuery(&paginationStateParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.ListStatesParams{
		Limit:  paginationStateParams.PageSize,
		Offset: (paginationStateParams.PageID - 1) * paginationStateParams.PageSize,
	}

	states, err := server.store.ListStates(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, states)
}

type deleteStateRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteState(ctx *gin.Context) {
	var req deleteStateRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := server.store.DeleteState(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "State deleted successfully",
	})
}

type updateStateIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateStateNameRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updateState(ctx *gin.Context) {
	var reqName updateStateNameRequest
	var reqId updateStateIDRequest

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

	_, err := server.store.GetState(ctx, reqId.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	args := database.UpdateStateParams{
		ID:   reqId.ID,
		Name: reqName.Name,
	}

	updatedState, err := server.store.UpdateState(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedState)

}
