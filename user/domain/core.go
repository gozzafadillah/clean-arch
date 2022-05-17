package UserDomain

import "time"

type Users struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Role      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
