package user

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(data Core) error
	CheckByEmail(email string) (*Core, error)
}

type UseCaseInterface interface {
	Create(data Core) error
}
