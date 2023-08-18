package api

import (
	"database/sql"
	"log"
	"net/http"

	database "github.com/fajardodiaz/infosgroup-employee-management/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type createPositionReq struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) createPosition(ctx *gin.Context) {
	var req createPositionReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	position, err := server.store.CreatePosition(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}

type getPositionByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getPositionById(ctx *gin.Context) {
	var req getPositionByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	position, err := server.store.GetPosition(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, position)
}

var paginationPositionParams = struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=20"`
}{
	PageID:   1,
	PageSize: 20,
}

func (server *Server) getPositions(ctx *gin.Context) {
	if err := ctx.ShouldBindQuery(&paginationPositionParams); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := database.ListPositionsParams{
		Limit:  paginationPositionParams.PageSize,
		Offset: (paginationPositionParams.PageID - 1) * paginationPositionParams.PageSize,
	}

	positions, err := server.store.ListPositions(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, positions)
}

type deletePositionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deletePosition(ctx *gin.Context) {
	var req deletePositionRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := server.store.DeletePosition(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Position deleted successfully",
	})
}

type updatePositionIDRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updatePositionNameRequest struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) updatePosition(ctx *gin.Context) {
	var reqName updatePositionNameRequest
	var reqId updatePositionIDRequest

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

	_, err := server.store.GetPosition(ctx, reqId.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	args := database.UpdatePositionParams{
		ID:   reqId.ID,
		Name: reqName.Name,
	}

	updatedPosition, err := server.store.UpdatePosition(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, updatedPosition)

}
