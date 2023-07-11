package dto

type LoginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	Token   string `json:"token"`
	Profile string `json:"profile"`
}

type LoginErrorResponse BaseResponse
