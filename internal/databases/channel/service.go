package channel

import "netradio/models"

type Service interface {
	GetAll() []models.ChannelInfo
	GetInfo(channelID int) models.ChannelInfo
}

func NewService() *service {
	return &service{}
}

type service struct{}

func (s service) GetAll() []models.ChannelInfo {
	return nil
}

func (s service) GetInfo(channelID int) models.ChannelInfo {
	return models.ChannelInfo{}
}
