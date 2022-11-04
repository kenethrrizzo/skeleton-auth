package users

type UserRequest struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
