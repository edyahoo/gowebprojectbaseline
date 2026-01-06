package app

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Authenticate(email, password string) error {
	return nil
}
