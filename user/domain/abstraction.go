package userDomain

type Service interface {
	// public
	Login(username, password string) (string, error)
	InsertData(domain Users) (response Users, err error)
	GetId(id int) (response Users, err error)
	GetUsername(username string) (Users, error)

	// Admin
	BanUser(username string) (Users, error)
}

type Repository interface {
	Save(domain Users) (id int, err error)
	GetUsernamePassword(username, password string) (domain Users, err error)
	GetById(id int) (domain Users, err error)
	GetByUsername(username string) (int, error)
	BanUser(id int, domain Users) (Users, error)
}
