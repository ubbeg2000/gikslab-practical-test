package dto

type RegistrationBody struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Profile  string   `json:"profile"`
	Skill    []string `json:"skill"`
}

type RegistrationResponse BaseResponse
