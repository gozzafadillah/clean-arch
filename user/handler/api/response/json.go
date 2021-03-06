package response

import userDomain "github.com/gozzafadillah/user/domain"

type ResponseJSON struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}

func FromDomain(domain userDomain.Users) ResponseJSON {
	return ResponseJSON{
		Name:     domain.Name,
		Username: domain.Username,
		Role:     domain.Role,
		Status:   domain.Status,
	}
}
