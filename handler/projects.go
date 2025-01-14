package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/model/db"
	"net/http"
	"strconv"
)

// Getprojects godoc
//
//	@Summary		Get one's projects
//	@Tags			projects
//	@Description	Get all the projects from user(login required)
//	@Produce		json
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	db.ProposalInfo
//	@Failure		404				{object}	handler.Response	"Resource requested not found"
//	@Router			/users/myproject [get]
func Getprojects(r *gin.Context) {
	id := r.GetInt("userID")
	data, n := model.GetManySth(db.ProposalInfo{UID: int32(id)})
	if n == 0 {
		SendError(r, model.ErrNotFound, nil, model.ErrorSender(), http.StatusNotFound)
		return
	}
	SendResponse(r, data)
}

// GetProject godoc
//
//	@Summary		Get a project
//	@Tags			projects
//	@Description	Get a project with its id
//	@Param			info_id			query	string	true	"the id of the project"
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.ProposalInfo
//	@Failure		403	{object}	handler.Response
//	@Router			/project [get]
func GetProject(r *gin.Context) {
	q := r.Query("info_id")
	id, err := strconv.Atoi(q)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	data := db.ProposalInfo{
		InfoID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, data)
}

// UpdateProject godoc
//
//	@Summary		Update one's project
//	@Tags			projects
//	@Description	Update user's project(login required)
//	@Accept			application/json
//	@Param			id				query	string			true	"the id of the project"
//	@Param			Authorization	header	string			true	"token"
//	@Param			newproject		body	db.ProposalInfo	true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		500	{object}	handler.Response
//	@Failure		403	{object}	handler.Response
//	@Failure		401	{object}	handler.Response
//	@Router			/project [put]
func UpdateProject(r *gin.Context) {
	sid := r.Query("id")
	if sid == "" {
		SendBadRequest(r, model.ErrBadRequest, nil, model.ErrorSender())
		return
	}
	id, _ := strconv.Atoi(sid)
	uid := r.GetInt("userID")
	data := db.ProposalInfo{
		InfoID: int32(id),
	}
	data = model.GetSth(data)
	if data.UID != int32(uid) {
		SendError(r, model.ErrNotAuthorized, nil, model.ErrorSender(), http.StatusUnauthorized)
		return
	}
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendBadRequest(r, model.ErrBadRequest, data, model.ErrorSender())
		return
	}
	err = model.UpdateSth(data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, NoResponse)
}

// CreateProject godoc
//
//	@Summary		Create a new project
//	@Tags			projects
//	@Description	Create user's project(login required)
//	@Accept			application/json
//	@Param			Authorization	header	string			true	"token"
//	@Param			newproject		body	db.ProposalInfo	true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Router			/project/newproject [post]
func CreateProject(r *gin.Context) {
	data := new(db.ProposalInfo)
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	data.UID = int32(r.GetInt("userID"))
	if data.Budget == "" {
		data.Budget = "{}"
	}
	if data.Nodes == "" {
		data.Nodes = "{}"
	}
	err, ret := model.CreateSth(*data)
	if err != nil {
		SendError(r, err, ret, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, ret)
}

// DeleteProject godoc
//
//	@Summary		Delete one's project
//	@Tags			projects
//	@Description	Delete user's project(login required)
//	@Accept			application/json
//	@Param			id				query	string	true	"the id of the project"
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		500	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Router			/project [delete]
func DeleteProject(r *gin.Context) {
	sid := r.Query("id")
	if sid == "" {
		SendBadRequest(r, model.ErrBadRequest, nil, model.ErrorSender())
		return
	}
	id, _ := strconv.Atoi(sid)
	err := model.DeleteProposal(db.ProposalInfo{
		InfoID: int32(id),
	})
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, NoResponse)
}
