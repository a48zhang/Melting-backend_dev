package handler

import (
	"main/model"
	"main/model/db"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginResponse struct {
	Code  int
	ID    int
	Token string
}

type loginRequest struct {
	NickName string `json:"nick_name"`
	QQ       string `json:"qq"`
	Auth     string `json:"auth"`
}

// Login godoc
//
//	@Summary		login
//	@Tags			register and login
//	@Description	login and return id&token
//	@Accept			application/json
//	@Param			loginAuth	body	model.loginRequest	true	"the User who is logging in"
//	@Param			loginType	query	string				false	"type of login(use 'qq' to login with qq)"
//	@Produce		json
//	@Success		200	{object}	model.loginResponse
//	@Failure		401	{object}	handler.Response	"username or password incorrect"
//	@Failure		403	{object}	handler.Response	"param not satisfied"
//	@Failure		500	{object}	handler.Response	"token generation failed"
//	@Router			/login [post]
func Login(r *gin.Context) {
	typ := r.Query("type")

	switch typ {
	case "MuxiPass":
		//TODO
	case "ccnu":
		//TODO
	case "qq":
		QQLogin(r)
	default:
		NativeLogin(r)
	}
}

func NativeLogin(r *gin.Context) {
	loginAuth := db.User{}
	err := r.ShouldBindJSON(&loginAuth)
	if err != nil {
		SendError(r, model.ErrBadRequest, loginAuth,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	if loginAuth.Auth == "" || loginAuth.NickName == "" {
		SendError(r, model.ErrAuthInvalid, loginAuth,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	token, err := service.LoginNative(loginAuth)
	if err == model.ErrAuthIncorrect {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusUnauthorized)
		return
	}
	if err != nil {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, loginResponse{
		Code:  http.StatusAccepted,
		ID:    int(loginAuth.UID),
		Token: token,
	})
}

func QQLogin(r *gin.Context) {
	loginAuth := db.User{}
	err := r.ShouldBindJSON(&loginAuth)
	if err != nil {
		SendError(r, model.ErrBadRequest, loginAuth,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	if loginAuth.Auth == "" || loginAuth.Qq == "" {
		SendError(r, model.ErrAuthInvalid, loginAuth,
			model.ErrorSender(), http.StatusBadRequest)
		return
	}
	token, err := service.LoginWithQQ(loginAuth)
	if err == model.ErrAuthIncorrect {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusUnauthorized)
		return
	}
	if err != nil {
		SendError(r, err, nil,
			model.ErrorSender(), http.StatusInternalServerError)
		return
	}
	SendResponse(r, loginResponse{
		Code:  http.StatusAccepted,
		ID:    int(loginAuth.UID),
		Token: token,
	})
}
