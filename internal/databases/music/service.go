package music

import (
	"netradio/models"
	"os"
	"time"
)

type Service interface {
	GetPodcasts() []models.MusicInfo
	LoadMusic(info models.MusicInfo, dst chan<- []byte) error
}

func NewService() *service {
	return &service{}
}

type service struct{}

func (s service) GetPodcasts() []models.MusicInfo {
	return nil
}

func (s service) LoadMusic(info models.MusicInfo, dst chan<- []byte) error {
	fileBytes, err := os.ReadFile("./Rammstein_DU_HAST.mp3")
	if err != nil {
		return err
	}
	for len(fileBytes) != 0 {
		msgSize := 1024
		if len(fileBytes) < msgSize {
			msgSize = len(fileBytes)
		}

		dst <- fileBytes[:msgSize]
		fileBytes = fileBytes[msgSize:]
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}
