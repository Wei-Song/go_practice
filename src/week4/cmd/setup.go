package main

import (
	"context"
	"go_practice/go_practice/src/week4/internal/repository"
	"go_practice/go_practice/src/week4/internal/repository/ent"
	"go_practice/go_practice/src/week4/internal/service"
	"go_practice/go_practice/src/week4/internal/usecase"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var UserSet = wire.NewSet(
	service.NewUserService,
	repository.NewRepository,
	usecase.NewUserUsecase,
)

// NewDB 初始化 db 连接
func NewDB(v *viper.Viper) (*ent.Client, error) {
	// fmt.Println(v.Sub("db").GetString("type"), v.Sub("db").GetString("dsn"))
	client, err := ent.Open(
		v.Sub("db").GetString("type"),
		v.Sub("db").GetString("dsn"),
	)
	if err != nil {
		return nil, err
	}

	// 数据迁移
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

// InitConfig 初始化配置文件
func InitConfig() (*viper.Viper, error) {
	viper.AddConfigPath("../config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}
