package service

import (
	"gin-server/config"
	"gin-server/entity"
)

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
	UpdateOne(int, entity.Video) entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	config.DB.Create(&video)
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	videos := []entity.Video{}
	config.DB.Find(&videos)
	return videos
}

func (service *videoService) UpdateOne(id int, video entity.Video) entity.Video {
	config.DB.Where("id = ?", id).First(&video)
	config.DB.Save(video)
	return video
}
