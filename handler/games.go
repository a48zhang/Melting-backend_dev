package handler

import (
	"main/model"
	"main/model/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GameSelect godoc
//
//	@Summary		Get a game's info
//	@Tags			projects
//	@Description	Get a game's info by id
//	@Accept			application/json
//	@Param			Authorization	header	string	true	"token"
//	@Param			game_id			query	string	true	"game_id"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Router			/project/games [get]
func GameSelect(r *gin.Context) {
	id, err := strconv.Atoi(r.Query("game_id"))
	if err != nil || id == 0 {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	data := new(db.Game)
	data.Gameid = int32(id)
	SendResponse(r, model.GetSth(*data))
}

// FindGames  godoc
//
//	@Summary		Get some games' info
//	@Tags			projects
//	@Description	Get some games' info with certain circumstances
//	@Accept			application/json
//	@Param			Authorization	header	string	true	"token"
//	@Param			data			body	db.Game	true	"game circumstances"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Router			/project/games/find [post]
func FindGames(r *gin.Context) {
	data := new(db.Game)
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	SendResponse(r, model.GetGames(*data))
}

// GameDetail godoc
//
//	@Summary		Get a game's detail
//	@Tags			projects
//	@Description	Get a game's detail by id
//	@Accept			application/json
//	@Param			Authorization	header	string	true	"token"
//	@Param			game_id			query	string	true	"game_id"
//	@Success		200				
//	@Failure		400				{object}	handler.Response
//	@Failure		404				
//	@Router			/project/games/details [get]
func GameDetail(r *gin.Context) {
	id := r.Query("game_id")
	if id == "" {
		SendError(r, nil, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	r.Redirect(http.StatusMovedPermanently, "/resource/games/"+id+".html")
}
