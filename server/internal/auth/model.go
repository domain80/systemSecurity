package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
  ID        int    `json:"id" db:"id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
	Role      string `json:"role" db:"role"`
	Password  string `json:"-" db:"password"` // Exclude Password from JSON output
	Email     string `json:"email" db:"email"`
}

type Log struct {
  Id string
  Time time.Time
  Action string
  UserID string
}

type Claims struct {
	ID        int     `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

