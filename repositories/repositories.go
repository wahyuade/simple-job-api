package repositories

type Repositories struct {
	UserRespository
}

func NewRepositories() Repositories {
	uR := newUserRepository()
	return Repositories{
		UserRespository: uR,
	}
}
