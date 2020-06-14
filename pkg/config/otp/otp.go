package otp

import "github.com/xlzd/gotp"

type otpService struct {
}

func NewOTPService() otpService{
	return otpService{}
}

func(o otpService) GenerateOTP() string{
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	return totp.Now()
}