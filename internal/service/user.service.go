package service

import "ecommerce-backend-golang/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRep(),
	}
}

func (us *UserService) GetListUser() string {
	return us.userRepo.GetListUser()
}