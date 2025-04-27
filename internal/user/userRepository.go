package user

type UserRepository interface {
	GetOneById(id string) (*User, error)
	GetAll() ([]User, error)
	Create(params UserCreateParams) (*User, error)
	Update(id string, params UserCreateParams) (*User, error)
	Delete(id string) error
}