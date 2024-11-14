package Interface

import "wan-api-verify-user/Model"

type IUserData interface {
	GetUserByUsername(username string) (*Model.KOL, error)
}