package structure

import (
	"github.com/hossein1376/kala/pkg/Password"
)

type User struct {
	ID        int               `json:"id"`
	FirstName string            `json:"firstname"`
	LastName  string            `json:"lastname"`
	Username  string            `json:"username"`
	Password  Password.Password `json:"-"`
	Email     string            `json:"email"`
	Phone     string            `json:"phone"`
	Role      string            `json:"role"`
	Orders    []Order           `json:"orders"`
}

type UserRequest struct {
	UserName  string  `json:"username"`
	Password  string  `json:"password"`
	FirstName *string `json:"firstname"`
	LastName  *string `json:"lastname"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type UserUpdateRequest struct {
	UserName  *string `json:"username"`
	Password  *string `json:"password"`
	FirstName *string `json:"firstname"`
	LastName  *string `json:"lastname"`
	Email     *string `json:"email"`
	Phone     *string `json:"phone"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  []byte `json:"-"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
}
