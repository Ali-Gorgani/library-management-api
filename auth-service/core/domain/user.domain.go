package domain

import "time"

type User struct {
	ID        uint    
	Username  string 
	Password  string
	Email     string 
	IsAdmin   bool   
	CreatedAt time.Time 
}
