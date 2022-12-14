package service

import(
	"github.com/romanSeleznev/markList/repository"
)

type Authorization interface{

}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}