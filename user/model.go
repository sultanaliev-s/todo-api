package user

import (
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	Image     *string
	CreatedOn time.Time
}
