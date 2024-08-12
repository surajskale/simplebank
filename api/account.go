package api

import (
	"database/sql"
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // how does go compiler know errorResponse is declared in server.go -> may be bcoz they are in same package. 	But what about context?
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int64 `form:"page_id" binding:"required,min=1"`          // min=1 tells that id should be >=1 without this it will send the error as invalid id
	PageSize int64 `form:"page_size" binding:"required,min=5,max=10"` // page size should be >= 5 && <= 10
}

func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // how does go compiler know errorResponse is declared in server.go -> may be bcoz they are in same package. 	But what about context?
		return
	}

	arg := db.ListAccountsParams{
		Limit:  int32(req.PageSize),
		Offset: int32(req.PageID-1) * int32(req.PageSize),
	}

	account, err := server.store.ListAccounts(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)

}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"` // min=1 tells that id should be >=1 without this it will send the error as invalid id
}

// Function to get specific account
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // how does go compiler know errorResponse is declared in server.go -> may be bcoz they are in same package. 	But what about context?
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)

}
