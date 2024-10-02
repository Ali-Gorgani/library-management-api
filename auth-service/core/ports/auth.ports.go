package ports

import "time"

type AuthRepository interface {
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
