package user

import (
	"time"
)

type User struct {
	CreatedAt time.Time
	ID        string
	Name      string
	Email     string
	Role      string
}
