package myservice

type UserRepository interface {
	Save() error
}

type User struct {
}

func (u *User) Save() error {
	return nil
}

func NewUserRepository() *User {
	return &User{}
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}
