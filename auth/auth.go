package auth

import (
	"fmt"
	"time"

	"github.com/arshabbir/utils/config"
	"github.com/arshabbir/utils/logger"
	"github.com/arshabbir/utils/model"
)

type auth struct {
	conf *config.Config
	l    logger.Logger
}

type Auth interface {
	Login(model.LogRequest) (*model.LoginResponse, *model.ApiError)
	VerifyToken(model.VerifyTokenRequest) (*model.VerifyTokenResponse, *model.ApiError)
}

func NewAuth(conf *config.Config, l logger.Logger) Auth {
	return &auth{conf: conf, l: l}
}

func (a *auth) Login(l model.LogRequest) (*model.LoginResponse, *model.ApiError) {
	// Implement the REST api client to interact with auth service and get the token
	a.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "Auth", Message: fmt.Sprintf("Login reuest from : %v", l)})
	return nil, nil
}

func (a *auth) VerifyToken(t model.VerifyTokenRequest) (*model.VerifyTokenResponse, *model.ApiError) {
	// Implement the REST api client to interact with auth service and get the token
	a.l.Log(model.LogRequest{Timestamp: time.Now(), ServiceName: "Auth", Message: fmt.Sprintf("Token reuest  : %v", t)})
	return nil, nil
}
