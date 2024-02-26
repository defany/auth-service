package model

const (
	LogCreateUser = "create_user"
	LogDeleteUser = "delete_user"
	LogUpdateUser = "update_user"
	LogGetUser    = "get_user"
)

type Log struct {
	Action string
	UserID uint64
}
