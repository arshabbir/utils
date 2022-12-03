package auth

import (
	"utilsmod/config"
	"utilsmod/model"
)

type auth struct {
	conf config.Config
}

type Auth interface {
	Login(model.LogRequest) (*model.LoginResponse, *model.ApiError)
	VerifyToken(model.VerifyTokenRequest) (*model.VerifyTokenResponse, *model.ApiError)
}

func NewAuth(conf config.Config) Auth {
	return &auth{conf: conf}
}

func (a *auth) Login(l model.LogRequest) (*model.LoginResponse, *model.ApiError) {
	// Implement the REST api client to interact with auth service and get the token
	return nil, nil
}

func (a *auth) VerifyToken(t model.VerifyTokenRequest) (*model.VerifyTokenResponse, *model.ApiError) {
	// Implement the REST api client to interact with auth service and get the token
	return nil, nil
}
