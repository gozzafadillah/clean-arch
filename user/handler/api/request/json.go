package request

import UserDomain "github.com/gozzafadillah/user/domain"

type RequestJSON struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToDomain(req RequestJSON) UserDomain.Users {
	return UserDomain.Users{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}
