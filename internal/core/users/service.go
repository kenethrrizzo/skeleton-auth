package users

type Service interface {
	SignUp(user *User) (*User, error)
	SignIn(user *User) (*User, error)
}

type UserService struct {
	userRepo map[string]string
}

func NewUserService(userRepo map[string]string) Service {
	return &UserService{userRepo}
}

func (us *UserService) SignUp(user *User) (*User, error) {
	return nil, nil
}

func (us *UserService) SignIn(user *User) (*User, error) {
	return nil, nil
}