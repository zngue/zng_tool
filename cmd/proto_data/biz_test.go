package main

import "context"

type UserNameRepo interface {
}
type UserNameUseCase struct {
	UserNameRepo UserNameRepo
}

func NewUserNameUseCase(UserNameRepo UserNameRepo) *UserNameUseCase {
	return &UserNameUseCase{
		UserNameRepo: UserNameRepo,
	}
}
func (uc *UserNameUseCase) Info(ctx context.Context, req any) (rs any, err error) {

	return
}
