package otp

type IOtp interface {
	GenerateOTP() string
}
