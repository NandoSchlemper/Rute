package entities

type IClient interface{}

type Client struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
