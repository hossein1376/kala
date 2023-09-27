package structure

import (
	"github.com/hossein1376/kala/pkg/Password"
)

type User struct {
	ID        int               `json:"id"`
	Username  string            `json:"username"`
	Password  Password.Password `json:"-"`
	FirstName string            `json:"firstname,omitempty"`
	LastName  string            `json:"lastname,omitempty"`
	Email     string            `json:"email,omitempty"`
	Phone     string            `json:"phone,omitempty"`
	Orders    []Order           `json:"orders,omitempty"`
	Role      string            `json:"-"`
}

type UserCreationRequest struct {
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
