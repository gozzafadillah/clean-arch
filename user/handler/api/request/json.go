package request

import UserDomain "github.com/gozzafadillah/user/domain"

type RequestJSON struct {
	ID       int
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ToDomain(req RequestJSON) UserDomain.Users {
	return UserDomain.Users{
		ID:       req.ID,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}
}
