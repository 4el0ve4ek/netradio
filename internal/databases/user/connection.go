package user

import (
	"errors"
	"netradio/models"
)

type Service interface {
	GetUser(email, password string) (models.User, error)
	GetUserByUID(uid int) models.User
	AddUser(email, password string) (models.User, error)
	ChangeUserNickname(uid int, name string) error
	ChangePassword(uid int, password string) error
	ChangeEmail(uid int, email string) error
	Delete(uid int) error
}

func NewService() *databaseService {
	return &databaseService{}
}

type databaseService struct {
	uids  int
	users []models.User
}

func (d *databaseService) GetUser(email, password string) (models.User, error) {
	for _, user := range d.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return models.User{}, errors.New("no such user")
}

func (d *databaseService) AddUser(email, password string) (models.User, error) {
	for _, user := range d.users {
		if user.Email == email {
			return models.User{}, errors.New("email already used")
		}
	}
	d.uids++
	return models.User{
		UID:       d.uids,
		Nickname:  email,
		PhotoLink: "",
		Lang:      models.Ru,
		Status:    models.UserRegistered,
		Email:     email,
		Password:  password,
	}, nil
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
	for i, user := range d.users {
		if user.UID == uid {
			d.users[i].Nickname = name
			return nil
		}
	}
	return errors.New("no such user")
}

func (d *databaseService) ChangePassword(uid int, password string) error {
	for i, user := range d.users {
		if user.UID == uid {
			d.users[i].Password = password
			return nil
		}
	}
	return errors.New("no such user")
}

func (d *databaseService) ChangeEmail(uid int, email string) error {
	for i, user := range d.users {
		if user.UID == uid {
			d.users[i].Email = email
			return nil
		}
	}
	return errors.New("no such user")
}

func (d *databaseService) Delete(uid int) error {
	for i, user := range d.users {
		if user.UID == uid {
			d.users = append(d.users[:i], d.users[i+1:]...)
			return nil
		}
	}
	return errors.New("no such user")
}
