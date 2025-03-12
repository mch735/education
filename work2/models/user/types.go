package user

import (
	"time"
)

type User struct {
	CreatedAt time.Time
	Name      string
	Email     string
	Role      string
	ID        int
}
