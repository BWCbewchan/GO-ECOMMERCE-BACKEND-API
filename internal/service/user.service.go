package service

import "github.com/BWCbewchan/GO-ECOMMERCE-BACKEND-API/internal/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserRepo() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}

// user repo : us
func (us *UserService) GetInforUser() string {
	return us.userRepo.GetInforUser()
}
