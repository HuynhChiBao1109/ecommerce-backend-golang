package service

import "ecommerce-backend-golang/internal/repo"

// type UserService struct {
// 	userRepo *repo.UserRepo
// }

// func NewUserService() *UserService {
// 	return &UserService{
// 		userRepo: repo.NewUserRep(),
// 	}
// }

// func (us *UserService) GetListUser() string {
// 	return us.userRepo.GetListUser()
// }

// Apply interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo *repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: &userRepo,
	}
}

// Register implements IUserService.
func (u *userService) Register(email string, purpose string) int {
	panic("unimplemented")
}
