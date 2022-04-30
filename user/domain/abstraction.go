package UserDomain

type Service interface {
	Login(username, password string) (string, error)
	InsertData(domain Users) (response Users, err error)
}

type Repository interface {
	Save(domain Users) (id int, err error)
	GetUsernamePassword(username, password string) (domain Users, err error)
	GetById(id int) (domain Users, err error)
}
