package news

import (
	"errors"
	"netradio/models"
	"strconv"
	"time"
)

type Service interface {
	GetAll() []models.News
	GetByID(id int) (models.News, error)
	Add(news models.News) int
}

func NewService() *databaseService {
	return &databaseService{
		news: map[int]models.News{
			1: {
				Title:           "first",
				PublicationTime: time.Date(2013, 10, 2, 12, 14, 15, 0, time.Local).Unix(),
				Content:         "",
			},
			2: {
				Title:           "second",
				PublicationTime: time.Date(2013, 10, 2, 12, 14, 15, 0, time.Local).Unix(),
				Content:         "content",
			},
		},
	}
}

type databaseService struct {
	news map[int]models.News
}

func (ds *databaseService) GetAll() []models.News {
	res := make([]models.News, 0, len(ds.news))
	for _, newsOne := range ds.news {
		res = append(res, newsOne)
	}
	return res
}

func (ds *databaseService) GetByID(id int) (models.News, error) {
	newsOne, exists := ds.news[id]
	if !exists {
		return models.News{}, errors.New("no news for id=" + strconv.Itoa(id))
	}
	return newsOne, nil
}

var ids = 2

func (ds *databaseService) Add(model models.News) int {
	if model.ID == 0 {
		model.ID = ids
		ids++
	}
	ds.news[model.ID] = model
	return model.ID
}
