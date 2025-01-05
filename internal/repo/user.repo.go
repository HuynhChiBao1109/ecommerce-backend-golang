package repo

// type UserRepo struct{}

// func NewUserRep() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetListUser() string {
// 	return `Get list user`
// }

// Interface -version

// Apply interface
type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct{}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}

// GetUserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) bool {
	panic("unimplemented")
}
