package xui

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/mhshahin/xray-tg/pkg/config"
	"github.com/mhshahin/xray-tg/pkg/xui/models"
)

type XUIInterface interface {
	Login() error
	CreateUser() (*models.User, error)
	UpdateUser() (*models.User, error)
	DeleteUser() error
	RestartService() error
}

type XUI struct {
	restyClient *resty.Client
}

func NewXUI() XUIInterface {
	restyClient := resty.New()
	restyClient.SetTimeout(config.Cfg.XUI.Timeout)

	return &XUI{
		restyClient: restyClient,
	}
}

func (x *XUI) Login() error {
	url := fmt.Sprintf("%s/login", config.Cfg.XUI.Address)

	resp, err := x.restyClient.R().
		SetBody(&models.LoginRequest{
			Username: config.Cfg.XUI.Username,
			Password: config.Cfg.XUI.Password,
		}).Post(url)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Not OK")
	}

	cookies := []*http.Cookie{
		{
			Name:  "lang",
			Value: "en-US",
		},
		{
			Name:  "session",
			Value: resp.Header().Values("Set-Cookie")[0],
		},
	}

	x.restyClient.SetCookies(cookies)

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
