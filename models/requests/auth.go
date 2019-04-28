package requests

type Register struct {
	Email      string  `json:"email"`
	Name       string  `json:"name"`
	Password   string  `json:"password"`
}

type Login struct {
	Email      string  `json:"email"`
	Password   string  `json:"password"`
}