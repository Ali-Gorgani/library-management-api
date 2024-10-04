package repository

import "time"

type UserRes struct {
	ID       int    
	Username string 
	Email    string 
	IsAdmin  bool   
	CreatedAt time.Time
}
