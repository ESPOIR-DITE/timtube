package user

import (
	"errors"
	"timtube/api"
	"timtube/domain"
)

const videoDataURL = api.BASE_URL + "video/video-data/"

func CreateVideoData(use domain.VideoData) (domain.VideoData, error) {
	entity := domain.VideoData{}
	resp, _ := api.Rest().SetBody(use).Post(videoDataURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateVideoData(role domain.VideoData) (domain.VideoData, error) {
	entity := domain.VideoData{}
	resp, _ := api.Rest().SetBody(role).Post(videoDataURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil

}
func ReadVideoRawData(id string) (domain.VideoData, error) {
	entity := domain.VideoData{}
	resp, _ := api.Rest().Get(videoDataURL + "getRwa/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideoData(id string) (domain.VideoData, error) {
	entity := domain.VideoData{}
	resp, _ := api.Rest().Get(videoDataURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideoPicture(id string) ([]byte, error) {
	entity := []byte{}
	resp, _ := api.Rest().Get(videoDataURL + "video-picture/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func GetVideoPicture(id string) ([]byte, error) {
	entity := []byte{}
	resp, _ := api.Rest().Get(videoDataURL + "get/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteVideoData(id string) (domain.VideoData, error) {
	entity := domain.VideoData{}
	resp, _ := api.Rest().Get(videoDataURL + "delete/" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllPublicVideoData() ([]domain.VideoVideoData, error) {
	entity := []domain.VideoVideoData{}
	resp, _ := api.Rest().Get(videoDataURL + "video-public-picture")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadVideoDatas() ([]domain.VideoData, error) {
	entity := []domain.VideoData{}
	resp, _ := api.Rest().Get(videoDataURL + "getAll")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
