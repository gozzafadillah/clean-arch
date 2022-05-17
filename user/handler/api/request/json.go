package request

import (
	"time"

	userDomain "github.com/gozzafadillah/user/domain"
)

type RequestJSON struct {
	ID        int
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToDomain(req RequestJSON) userDomain.Users {
	return userDomain.Users{
		ID:       req.ID,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Status:   req.Status,
		Role:     req.Role,
	}
}
