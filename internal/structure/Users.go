package structure

import (
	"github.com/hossein1376/kala/pkg/Password"
)

type User struct {
	ID        int               `json:"id" example:"1"`
	Username  string            `json:"username" example:"user"`
	Password  Password.Password `json:"-"`
	FirstName string            `json:"firstname,omitempty" example:"John"`
	LastName  string            `json:"lastname,omitempty" example:"Doe"`
	Email     string            `json:"email,omitempty" example:"email@wmail.com"`
	Phone     string            `json:"phone,omitempty" example:"0123456789"`
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
