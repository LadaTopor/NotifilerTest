package service

type Note struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type RegisterUser struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"jwt_token"`
}
type AuthUser struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"jwt_token"`
}

type QuoteResponse struct {
	Quote struct {
		Body string `json:"body"`
	} `json:"quote"`
}
