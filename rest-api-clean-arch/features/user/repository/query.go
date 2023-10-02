package repository

import (
	"fakhry/clean-arch/features/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// CheckByEmail implements user.DataInterface.
func (*userRepository) CheckByEmail(email string) (*user.Core, error) {
	panic("unimplemented")
}

// Insert implements user.DataInterface.
func (repo *userRepository) Insert(data user.Core) error {
	//mapping dari struct core ke struct gorm/model
	var input = User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
	}
	tx := repo.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func New(db *gorm.DB) user.DataInterface {
	return &userRepository{
		db: db,
	}
}
