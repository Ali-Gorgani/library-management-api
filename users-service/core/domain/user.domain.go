package domain

type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	CreatedAt      string `json:"created_at"`
}

type AddUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UpdateUserParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UserLoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RolePermissions defines which actions each role can perform
var RolePermissions = map[string][]string{
    "User":  {"GetBooks", "AddBook", "BorrowBook", "ReturnBook"},
    "Admin": {"UpdateBook", "DeleteBook", "GetBooks", "AddBook", "BorrowBook", "ReturnBook"},
}
