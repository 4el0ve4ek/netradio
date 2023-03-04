package music

import (
	"bufio"
	"io"
	"log"
	"netradio/models"
	"netradio/pkg/errors"
	"netradio/pkg/generics/slices"
	"os"
)

const (
	chunkSize = 1 << 20 // 512kb
)

type Service interface {
	GetPodcasts() []models.MusicInfo
	LoadMusicBatch(info models.MusicInfo) (<-chan []byte, error)
}

func NewService() *service {
	return &service{}
}

type service struct{}

func (s service) GetPodcasts() []models.MusicInfo {
	return nil
}

func (s service) LoadMusicBatch(info models.MusicInfo) (<-chan []byte, error) {
	file, err := os.Open("./music.mp3")
	if err != nil {
		return nil, errors.Wrap(err, "open file")
	}
	r := bufio.NewReader(file)
	chunks := make(chan []byte)
	go func() {
		defer close(chunks)

		buf := make([]byte, 0, chunkSize)
		for {
			n, err := r.Read(buf[:cap(buf)])
			buf = buf[:n]
			if n == 0 {
				if err == nil {
					continue
				}
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}

			chunks <- slices.Copy(buf)
		}
	}()
	return chunks, nil
}
