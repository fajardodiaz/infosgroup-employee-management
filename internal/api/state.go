package api

import (
	"database/sql"
	"net/http"

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

func (server *Server) GetStateById(ctx *gin.Context) {
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
