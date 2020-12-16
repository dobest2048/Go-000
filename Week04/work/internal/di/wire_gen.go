//go:generate wire
//+build !wireinject

package di

import (
	"github.com/jikebang/Go-000/week04/work/configs"
	"github.com/jikebang/Go-000/week04/work/internal/dao"
	"github.com/jikebang/Go-000/week04/work/internal/service"
	"github.com/jikebang/Go-000/week04/work/pkg/email"
)

func NewService(c *configs.DbConfig, m *configs.EMailConfig) (*service.Service, error) {
	mailSender := email.NewMailSender(m)
	sqlDB, err := configs.NewDb(c)
	if err != nil {
		return nil, err
	}
	userRepo := dao.NewUserRepo(sqlDB)
	serviceService := service.NewService(mailSender, userRepo)
	return serviceService, nil
}
