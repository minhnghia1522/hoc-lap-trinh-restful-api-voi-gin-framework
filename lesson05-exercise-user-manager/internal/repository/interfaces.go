package repository

type InMemoryUserRepository interface {
	SelectByCondition()
	SelectByUUID(uuid string)
	CreateUser()
	UpdateUser()
	DeleteUser(uuid string)
}
