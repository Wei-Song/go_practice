package main

import (
	"go_practice/go_practice/src/week4/internal/repository"
	"go_practice/go_practice/src/week4/internal/service"
	"go_practice/go_practice/src/week4/internal/usecase"
)

// Injectors from wire.go:

func InitUserService() (*service.UserService, error) {
	viper, err := InitConfig()
	if err != nil {
		return nil, err
	}
	client, err := NewDB(viper)
	if err != nil {
		return nil, err
	}
	iUserRepo := repository.NewRepository(client)
	iUserUsecase := usecase.NewUserUsecase(iUserRepo)
	userService := service.NewUserService(iUserUsecase)
	return userService, nil
}
