package user

import (
	"netradio/models"
	"netradio/pkg/log"
)

type Service interface {
	GetUser(email, password string) models.User
	AddUser(email, password string) models.User
	ChangeUserNickname(uid int, name string) error
	ChangePassword(uid int, password string) error
	ChangeEmail(uid int, email string) error
	CreateUser(email, password string, imageID int) error
	Delete(user models.User)
}

func NewService(logger log.Logger) *databaseService {
	return &databaseService{}
}

type databaseService struct{}

func (d *databaseService) GetUser(email, password string) models.User {
	if email == "test" {
		return d.GetUserByUID(123)
	}
	return models.User{}
}

func (d *databaseService) AddUser(email, password string) models.User {
	if email == "test" {
		return d.GetUserByUID(123)
	}
	return models.User{}
}

func (d *databaseService) GetUserByUID(uid int) models.User {
	return models.User{
		UID:       uid,
		Nickname:  "Harrley",
		PhotoLink: "https://mobimg.b-cdn.net/v3/fetch/00/009384f823da78adb2a30695e4502d5e.jpeg", /// think
		Lang:      models.Ru,
		Status:    models.UserRegistered,
	}
}

func (d *databaseService) ChangeUserNickname(uid int, name string) error {
	return nil
}

func (d *databaseService) ChangePassword(uid int, password string) error {
	return nil
}

func (d *databaseService) ChangeEmail(uid int, email string) error {
	return nil
}

func (d *databaseService) CreateUser(email, password string, imageID int) error {
	return nil
}

func (d *databaseService) Delete(user models.User) {}
