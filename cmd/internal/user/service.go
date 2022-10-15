package user

import "context"

type Service struct {
	storage Storage
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (User, error) {
	//TODO implement me
	return User{}, nil
}
