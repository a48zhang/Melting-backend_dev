package handler

import (
	"encoding/json"
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UploadProfile godoc
//
//	@Summary		upload profile
//	@Tags			user
//	@Description	upload user's profile
//	@Accept			application/json
//	@Produce		json
//	@Param			Authorization	header		string	true	"token"
//	@Param			Profile			body		db.User	true	"new user profile"
//	@Success		200				{string}	string
//	@Failure		500				{object}	handler.Response	"update failed"
//	@Router			/users [put]
func UploadProfile(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	err := r.ShouldBindJSON(&data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	if err := model.UpdateSth(data); err != nil {
		SendError(r, err, data, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, model.NoResponse)
}

// UploadPhoto godoc
//
//	@Summary		upload photo
//	@Tags			user
//	@Description	upload user's avatar
//	@Accept			image/jpeg
//	@Produce		json
//	@Param			file			formData	object	true	"the photo of the user"
//	@Param			Authorization	header		string	true	"token"
//	@Success		200				{object}	string
//	@Failure		400				{object}	handler.Response	"file not received"
//	@Failure		500				{object}	handler.Response	"failed to upload file"
//	@Router			/users/photo [put]
func UploadPhoto(r *gin.Context) {
	id := r.GetInt("userID")
	H, err := r.FormFile("file")
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	file, err := H.Open() // Warning: file must be *.jpg
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusBadRequest)
		return
	}
	url, err := service.UploadProfilePhoto(&file, H.Size)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	var data = db.User{
		UID: int32(id),
	}
	data = model.GetSth(data)
	data.Photo = url
	err = model.UpdateSth(data)
	if err != nil {
		SendError(r, err, nil, model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, data)
}

// GetUserInfo godoc
//
//	@Summary		Get User's info
//	@Tags			user
//	@Description	Get User's info
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Router			/users [get]
func GetUserInfo(r *gin.Context) {
	id := r.GetInt("userID")
	data := db.User{
		UID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, data)
}

// GetOnesInfo godoc
//
//	@Summary		Get User's info
//	@Tags			user
//	@Description	Get User's info with its userID
//	@Param			id	query	string	true	"uid"
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Router			/user [get]
func GetOnesInfo(r *gin.Context) {
	Q := r.Query("id")
	id, _ := strconv.Atoi(Q)
	data := db.User{
		UID: int32(id),
	}
	data = model.GetSth(data)
	SendResponse(r, data)
}

// JoinProposal godoc
//
//	@Summary		Join certain proposal (login required)
//	@Tags			user
//	@Description	Join a proposal with infoId
//	@Param			id				query	string	true	"infoId"
//	@Param			Authorization	header	string	true	"token"
//	@Produce		json
//	@Success		200	{object}	db.User
//	@Failure		400				{object}	handler.Response	"user exists"
//	@Router			/join [get]
func JoinProposal(r *gin.Context) {
	uid := r.GetInt("userID")
	infoID, _ := strconv.Atoi(r.Query("id"))

	data := db.ProposalInfo{
		InfoID: int32(infoID),
	}
	data = model.GetSth(data)
	corplist := make([]interface{}, 100)
	err := json.Unmarshal([]byte(data.Corporates), &corplist)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), 500)
		return
	}
	for i := range corplist {
		if i == uid {
			SendError(r, model.ErrUserExist, data, model.ErrorSender(), 400)
			return
		}
	}
	corplist = append(corplist, uid)
	tmp, _ := json.Marshal(corplist)
	data.Corporates = string(tmp)
	err = model.UpdateSth(data)
	if err != nil {
		SendError(r, err, data, model.ErrorSender(), 500)
		return
	}
	SendResponse(r, data)
}
