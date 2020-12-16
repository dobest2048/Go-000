package di

import (
	"github.com/google/wire"
	"github.com/jikebang/Go-000/week04/work/configs"
	"github.com/jikebang/Go-000/week04/work/internal/dao"
	"github.com/jikebang/Go-000/week04/work/internal/service"
	"github.com/jikebang/Go-000/week04/work/pkg/email"
)

// +build wireinject

// NewService 定义injector的函数签名
func NewService(c *configs.DbConfig, m *configs.EMailConfig) (*service.Service, error) {
	wire.Build(service.NewService, email.EMailSet, dao.UserSet, configs.NewDb)
	return &service.Service{}, nil
}
