package contract

import (
	"authentication/data"
)

type UserService interface {
	FindByEmail(email string) (user data.User, err error)
}

type UserRepository interface {
	FindByEmail(email string) (user data.User, err error)
}
