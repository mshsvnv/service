package dto

import "github.com/mshsvnv/service/model"

type LoginReq struct {
	Id       int
	Email    string
	Password string
}

type RegisterReq struct {
	Name     string
	Surname  string
	Email    string
	Password string
	Role     string
}

type UpdateRoleReq struct {
	Email string
	Role  model.UserRole
}
