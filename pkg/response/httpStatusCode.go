package response

const (
	ErrCodeSuccess      = 20001 //success
	ErrCodeParamInvalid = 20003 // Email is invalid

	ErrInvalidToken = 30003 //token is invalid
)

// mesage
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "token is invalid",
}
