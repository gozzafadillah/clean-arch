package response

import UserDomain "github.com/gozzafadillah/user/domain"

type ResponseJSON struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

func FromDomain(domain UserDomain.Users) ResponseJSON {
	return ResponseJSON{
		Name:     domain.Name,
		Username: domain.Username,
		Role:     domain.Role,
	}
}
