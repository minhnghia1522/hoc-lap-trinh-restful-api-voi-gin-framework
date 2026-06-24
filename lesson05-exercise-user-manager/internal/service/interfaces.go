package service

type UserService interface {
	Search()
	FindUserByUUID(uuid string)
	CreateUser()
	UpdateUser()
	DeleteUser(uuid string)
}
