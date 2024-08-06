package model

type VerifyOTP struct {
	Phone string `json:"username"`
	Otp   string `json:"otp"`
}
