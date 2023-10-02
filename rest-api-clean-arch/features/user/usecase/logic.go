package usecase

import (
	"errors"
	"fakhry/clean-arch/features/user"
)

type userUsecase struct {
	userRepository user.DataInterface
}

// Create implements user.UseCaseInterface.
func (uc *userUsecase) Create(data user.Core) error {
	//validasi
	if data.Email == "" || data.Password == "" {
		return errors.New("error. email dan password harus diisi")
	}

	err := uc.userRepository.Insert(data)
	return err
}

func New(userRepo user.DataInterface) user.UseCaseInterface {
	return &userUsecase{
		userRepository: userRepo,
	}
}
