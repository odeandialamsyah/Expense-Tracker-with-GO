package controllers

import "expense-tracker-with-go/repository"

type AuthController struct {
	UserRepo repository.UserRepository
	RoleRepo repository.RoleRepository
}