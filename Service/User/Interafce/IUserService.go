package Interface

import "wan-api-verify-user/Model"

type IUserService interface {
	GetUserByUsername(username string) (*Model.KOL, error)
}