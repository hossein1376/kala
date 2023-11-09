package structure

import (
	"github.com/hossein1376/kala/pkg/password"
)

type User struct {
	ID        int               `json:"id" example:"1"`
	Username  string            `json:"username" example:"user"`
	Password  password.Password `json:"-"`
	FirstName string            `json:"firstname,omitempty" example:"John"`
	LastName  string            `json:"lastname,omitempty" example:"Doe"`
	Email     string            `json:"email,omitempty" example:"email@email.com"`
	Phone     string            `json:"phone,omitempty" example:"0123456789"`
	Orders    []Order           `json:"orders,omitempty"`
	Role      string            `json:"-"`
}

type UserCreationRequest struct {
	UserName  string  `json:"username" example:"user"  validate:"required"`
	Password  string  `json:"password" example:"123456"  validate:"required"`
	FirstName *string `json:"firstname" example:"John"`
	LastName  *string `json:"lastname" example:"Doe"`
	Email     *string `json:"email" example:"email@email.com"`
	Phone     *string `json:"phone" example:"0123456789"`
}

type GetAllUsersRequest struct {
	Count int `schema:"count,required"`
	Page  int `schema:"page,required"`
}

type GetAllUsersResponse struct {
	Users       []*User `json:"users"`
	CurrentPage int     `json:"current_page" example:"2"`
	TotalPages  int     `json:"total_pages" example:"5"`
	PerPage     int     `json:"per_page" example:"10"`
}

type UserUpdateRequest struct {
	UserName  *string `json:"username" example:"user"`
	Password  *string `json:"password" example:"123456"`
	FirstName *string `json:"firstname" example:"John"`
	LastName  *string `json:"lastname" example:"Doe"`
	Email     *string `json:"email" example:"email@email.com"`
	Phone     *string `json:"phone" example:"0123456789"`
}
