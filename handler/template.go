package handler

import (
	"github.com/gin-gonic/gin"
	"main/model"
	"main/model/db"
	"net/http"
	"strconv"
)

// GetTemplate godoc
//
//	@Summary		Get a template
//	@Tags			template
//	@Description	Get a template with its id or name
//	@Param			id				query	string	false	"the id of the template"
//	@Param			name			query	string	false	"the name of the template"
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.Template
//	@Failure		403	{object}	handler.Response
//	@Failure		404	{object}	handler.Response
//	@Router			/project/template [get]
func GetTemplate(r *gin.Context) {
	name := r.Query("name")
	qid := r.Query("id")
	if name != "" {
		data := model.GetTemplate(name)
		if data.Context == "" {
			SendError(r, model.ErrNotFound, data, model.ErrorSender(), http.StatusNotFound)
			return
		}
		SendResponse(r, data)
		return
	} else if qid != "" {
		id, err := strconv.Atoi(qid)
		if err != nil || id == 0 {
			SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
			return
		}
		data := db.Template{Temid: int32(id)}
		data = model.GetSth(data)
		if data.Name == "" && data.Context == "" {
			SendError(r, model.ErrNotFound, data, model.ErrorSender(), http.StatusNotFound)
			return
		}
		SendResponse(r, data)
		return
	}
	SendError(r, model.ErrBadRequest, nil, model.ErrorSender(), http.StatusBadRequest)
}

// CreateTemplate godoc
//
//	@Summary		Create a new template
//	@Tags			template
//	@Description	Create a new template(login required)
//	@Accept			application/json
//	@Param			Authorization	header	string		true	"token"
//	@Param			newproject		body	db.Template	true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Failure		500	{object}	handler.Response
//	@Router			/project/template [post]
func CreateTemplate(r *gin.Context) {
	data := new(db.Template)
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	if data.Context == "" {
		(*data).Context = "{}"
	}
	err, id := model.CreateSth(*data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	(*data).Temid = int32(id)
	SendResponse(r, data)
}

// UpdateTemplate godoc
//
//	@Summary		Update certain template
//	@Tags			template
//	@Description	Create user's project(login required)
//	@Accept			application/json
//	@Param			Authorization	header	string		true	"token"
//	@Param			project			body	db.Template	true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Failure		500	{object}	handler.Response
//	@Router			/project/template [put]
func UpdateTemplate(r *gin.Context) {
	data := new(db.Template)
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	err = model.UpdateSth(*data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, data)
}

// DeleteTemplate godoc
//
//	@Summary		Delete certain template
//	@Tags			template
//	@Description	Create user's project(login required)
//	@Accept			application/json
//	@Param			Authorization	header	string	true	"token"
//	@Param			id				query	int		true	"operating project"
//	@Produce		json
//	@Success		200	{object}	handler.Response
//	@Failure		400	{object}	handler.Response
//	@Failure		500	{object}	handler.Response
//	@Router			/project/template [delete]
func DeleteTemplate(r *gin.Context) {
	id := r.Query("id")
	if id == "" {
		SendError(r, model.ErrBadRequest, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	atoi, _ := strconv.Atoi(id)
	data := db.Template{Temid: int32(atoi)}
	err := model.DeleteSth(data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, data)
}
