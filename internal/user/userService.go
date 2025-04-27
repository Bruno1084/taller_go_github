package user

type Service struct {
	repo UserRepository
}

func NewService(repo UserRepository) *Service {
	return &Service{repo: repo}
}

// Get one User by Id
func (s *Service) GetOneById(id string) (*User, error) {

}

// Get all Users
func (s *Service) GetAll() ([]User, error) {

} 

// Create one user
func (s *Service) Create(params UserCreateParams) (*User, error) {

}

// Update one user
func (s *Service) Update(id string, params UserCreateParams) (*User, error) {

}

// Delete one user
func (s *Service) Delete(id string) (*User, error) {
	
}

