package misc

type IdReturnStruct struct {
	Id int `json:"id"`
}

type TokenReturnStruct struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Id int `json:"id"`
	Password string `json:password`
}
