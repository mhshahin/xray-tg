package xui

import "github.com/mhshahin/xray-tg/pkg/xui/models"

type XUIInterface interface {
	Login() error
	CreateUser() (*models.User, error)
	UpdateUser() (*models.User, error)
	DeleteUser() error
	RestartService() error
}

type XUI struct{}

func NewXUI() XUIInterface {
	return &XUI{}
}

func (x *XUI) Login() error {
	return nil
}

func (x *XUI) ListUsers() {}

func (x *XUI) CreateUser() (*models.User, error) {
	return nil, nil
}

func (x *XUI) UpdateUser() (*models.User, error) {
	return nil, nil
}

func (x *XUI) DeleteUser() error {
	return nil
}

func (x *XUI) RestartService() error {
	return nil
}
