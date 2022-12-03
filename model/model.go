package model

import "time"

// E-com service modles
type Product struct {
	ProductID   string  `json:"productid"`
	ProductName string  `json:"productname"`
	Price       float64 `json:"price"`
}

type InventoryRequest struct {
	SeassonID  string   `json:"seassonid"`
	ProductIDs []string `json:"productids"`
}

type InventoryResponse struct {
	SeassonID string    `json:"seassonid"`
	Details   []Product `json:"details"`
}

type InventoryBook struct {
	SeassonID string    `json:"seassonid"`
	Products  []Product `json:"products"`
}

type PaymentRequest struct {
	SeassonID string    `json:"seassonid"`
	Method    string    `json:"method"`
	Products  []Product `json:"products"`
}

type PaymentResponse struct {
	SeassonID string   `json:"seassonid"`
	StatusVec []Status `json:"statusvec"`
}

type Status struct {
	Errorcode int    `json:"errorcode"`
	Cause     int    `json:"cause"`
	Message   string `json:"message"`
}

type ShipRequest struct {
	SeassonID    string    `json:"seassonid"`
	Products     []Product `json:"product"`
	ShippingAddr Address   `json:"address"`
}

type ShipResponse struct {
	SeassonID string   `json:"seassonid"`
	StatusVec []Status `json:"status"`
}

type Address struct {
	FlatNumber string `json:"flatnumber"`
	Street     string `json:"street"`
	LandMark   string `json:"landmark"`
	City       string `json:"city"`
	State      string `json:"state"`
	PinCode    string `json:"pincode"`
	Phone      string `json:"phone"`
}

// Auth models

type LoginRequest struct {
	AccessKey string `json:"accesskey"`
	SecretKey string `json:"secretkey"`
}

type LoginRespoinse struct {
	AccessKey string `json:"accesskey"`
	Token     string `json:"token"`
}

type VerifyTokenRequest struct {
	AccessKey string `json:"accesskey"`
	Token     string `json:"token"`
}

type VerifyTokenResponse struct {
	Token      string `json:"token"`
	Valid      bool   `json:"valid"`
	ExpiryTime string `json:"expirydate"`
}

// Logging models

type LogRequest struct {
	Timestamp   time.Time `json:"timestamp"`
	ServiceName string    `json:"servicename"`
	Message     string    `json:"message"`
}

// Tracing model

type TraceRequest struct {
}
