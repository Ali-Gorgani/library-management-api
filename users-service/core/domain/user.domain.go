package domain

type User struct {
	ID uint
	Username string
	HashedPassword string
	Email string
	CreatedAt string
}

type AddUserParam struct {
	Username string
	Password string
	Email string
}

type UpdateUserParam struct {
	Username string
	Password string
	Email string
}

type UserLoginParam struct {
	Username string
	Password string
}

